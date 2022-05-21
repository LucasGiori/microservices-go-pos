package model

import (
	"time"

	"github.com/google/uuid"
)

type TicketStatus string

const (
	OPENED TicketStatus = "OPENED"
	IN_PROGRESS TicketStatus = "IN_PROGRESS"
	CLOSED TicketStatus = "CLOSED"
)

type Ticket struct {
	Id          uuid.UUID
	OrderId     uuid.UUID
	Description string
	Email       string
	Status      TicketStatus `json:",omitempty"`
	DateTime    time.Time    `json:",omitempty"`
}
