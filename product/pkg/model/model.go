package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	Id    uuid.UUID `json:",omitempty"`
	Name  string
	Value *decimal.Decimal `json:",omitempty"`
}
