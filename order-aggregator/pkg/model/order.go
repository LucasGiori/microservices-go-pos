package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus int

const (
	OPEN OrderStatus = iota
)

func (s OrderStatus) String() string {
	return [...]string{"OPEN"}[s]
}

type Order struct {
	Id       uuid.UUID
	Customer *Customer
	DateTime *time.Time
	Items    []*OrderItem
	Status   OrderStatus
}

type OrderItem struct {
	Id        uuid.UUID
	Product   *Product
	Quantity  int32
	UnitValue decimal.Decimal
}
