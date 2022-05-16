package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	Id       uuid.UUID
	Customer *Customer
	DateTime time.Time
	Items    []*OrderItem
}

type OrderItem struct {
	Product   *Product
	Quantity  int32
	UnitValue decimal.Decimal
}
