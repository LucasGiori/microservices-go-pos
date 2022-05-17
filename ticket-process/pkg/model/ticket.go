package model

import (
	"time"

	"github.com/google/uuid"
)

type TicketStatus string

const (
	OPENED      TicketStatus = "OPENED"
	IN_PROGRESS              = "IN_PROGRESS"
	CLOSED                   = "CLOSED"
)

type Ticket struct {
	Id          uuid.UUID
	OrderId     uuid.UUID
	Description string
	Email       string
	Status      TicketStatus `json:",omitempty"`
	DateTime    time.Time    `json:",omitempty"`
}
