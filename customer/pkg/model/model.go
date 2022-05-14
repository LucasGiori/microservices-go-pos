package model

import "github.com/google/uuid"

type Customer struct {
	Id    uuid.UUID `json:",omitempty"`
	Name  string
	Email string
}
