package model

import "github.com/google/uuid"

type Customer struct {
	Id    uuid.UUID
	Name  string
	Email string
}
