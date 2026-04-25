package dtos

import "github.com/google/uuid"

type MessageDTO struct {
	Id       int64     `json:"id"`
	SenderId uuid.UUID `json:"senderId"`
	Subject  string    `json:"subject"`
	Body     string    `json:"body"`
	IsHtml   bool      `json:"isHtml"`
}
