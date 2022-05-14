package model

import "time"

type AuthRequest struct {
	Login    string
	Password string
}

type JWT struct {
	Token      string
	Expiration time.Time
}
