package model

type TicketStatus string

const (
	OPEN TicketStatus = "ABERTO"
)

type Ticket struct {
	Idproduct   string       `json:"id-product"`
	Description string       `json:"description"`
	Email       string       `json:"email"`
	Status      TicketStatus `json:"status"`
}
