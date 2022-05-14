package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:",omitempty"`
	Login    string
	Password string `json:",omitempty"`
}
