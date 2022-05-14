package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus int

const (
	CLOSED OrderStatus = iota
)

func (s OrderStatus) String() string {
	return [...]string{"CLOSED"}[s]
}

var OrderStatusMap = map[string]OrderStatus{"CLOSED": CLOSED}

type Order struct {
	Id       uuid.UUID `json:",omitempty"`
	Customer *Customer
	DateTime *time.Time
	Status   OrderStatus `json:",omitempty"`
	Items    []*OrderItem
	Total    decimal.Decimal `json:",omitempty"`
}

func (o *Order) CalcTotal() {
	total := decimal.Zero
	for _, i := range o.Items {
		i.CalcTotal()
		total = total.Add(i.Total)
	}

	o.Total = total
}

func (o Order) ValidateToCreate() error {
	if len(o.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	for _, i := range o.Items {
		if err := i.ValidateToCreate(); err != nil {
			return err
		}
	}

	return nil
}

type OrderItem struct {
	Id        uuid.UUID `json:"-"`
	Product   *Product
	Quantity  int32
	UnitValue decimal.Decimal
	Total     decimal.Decimal
}

func (i *OrderItem) CalcTotal() {
	i.Total = i.UnitValue.Mul(decimal.NewFromInt32((i.Quantity)))
}

func (i OrderItem) ValidateToCreate() error {
	if i.Quantity <= 0 {
		return errors.New("quantity must be more than 0")
	}

	if i.UnitValue.IsNegative() {
		return errors.New("item unit value must be a positve value")
	}

	return nil
}
