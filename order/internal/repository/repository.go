package repository

import (
	"context"
	"microservices/order/pkg/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

const (
	selectItemsByOrderId = "select oi.id, quantity, unit_value, total, oi.product_id, oi.product_name from order_item oi where order_id = $1;"
)

type Repository interface {
	Create(context.Context, *model.Order) (*model.Order, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	err := r.databaseManager.RunInTransaction(ctx, func(ctx context.Context) error {
		order, err := r.createOrder(ctx, order)
		if err != nil {
			return err
		}

		items, err := r.createOrderItems(ctx, order)
		if err != nil {
			return err
		}

		order.Items = items

		return nil
	})

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r RepositoryImpl) createOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	sql := "insert into \"order\" (customer_id, customer_name, date_time, status, total) values ($1, $2, $3, $4, $5) returning id;"
	var orderId uuid.UUID
	err := r.databaseManager.QueryRow(ctx, sql,
		order.Customer.Id, order.Customer.Name, order.DateTime, order.Status.String(), order.Total).Scan(&orderId)

	if err != nil {
		return nil, err
	}

	order.Id = orderId

	return order, nil
}

func (r RepositoryImpl) createOrderItems(ctx context.Context, order *model.Order) ([]*model.OrderItem, error) {
	queries := make([]database.BatchQuery, 0)
	sql := "insert into order_item (order_id, product_id, product_name, quantity, unit_value, total) values ($1, $2, $3, $4, $5, $6);"
	for _, i := range order.Items {
		queries = append(queries, database.BatchQuery{
			Sql: sql,
			Args: []interface{}{
				order.Id,
				i.Product.Id,
				i.Product.Name,
				i.Quantity,
				i.UnitValue,
				i.Total,
			},
		})
	}

	queries = append(queries, database.BatchQuery{
		Sql:  selectItemsByOrderId,
		Args: []interface{}{order.Id},
	})

	result, err := r.databaseManager.Batch(ctx, queries)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	for i := 0; i < len(order.Items); i++ {
		_, err = result.Exec()
		if err != nil {
			return nil, err
		}
	}

	rows, err := result.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return readOrderItems(rows)
}

func readOrderItems(rows pgx.Rows) ([]*model.OrderItem, error) {
	items := make([]*model.OrderItem, 0)

	for rows.Next() {
		var id uuid.UUID
		var quantity int32
		var unitValue decimal.Decimal
		var total decimal.Decimal
		var productId uuid.UUID
		var productName string

		if err := rows.Scan(&id, &quantity, &unitValue, &total, &productId, &productName); err != nil {
			return nil, err
		}

		items = append(items, &model.OrderItem{
			Id:        id,
			Quantity:  quantity,
			UnitValue: unitValue,
			Total:     total,
			Product: &model.Product{
				Id:   productId,
				Name: productName,
			},
		})
	}

	return items, nil
}
