package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id          uuid.UUID
	OrderId     string
	Description string
	Email       string
	Status      string
	DateTime    *time.Time `json:",omitempty"`
}
