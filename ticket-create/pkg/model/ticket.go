package model

import "github.com/google/uuid"

type TicketStatus string

const (
	OPEN TicketStatus = "OPENED"
)

type Ticket struct {
	Id          uuid.UUID
	OrderId     uuid.UUID
	Description string
	Email       string
	Status      TicketStatus
}
