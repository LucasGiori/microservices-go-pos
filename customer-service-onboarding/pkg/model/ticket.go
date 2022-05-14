package model

import "github.com/google/uuid"

type Ticket struct {
	Id          uuid.UUID
	Order       *Order
	Description string
}
