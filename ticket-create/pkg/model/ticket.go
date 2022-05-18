package model

import (
	"github.com/google/uuid"
)

type Ticket struct {
	Id          uuid.UUID
	Idproduct   string
	description string
	email       string
	status      string
}
