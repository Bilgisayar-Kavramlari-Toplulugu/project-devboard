package dtos

import (
	"time"

	"github.com/google/uuid"
)

type MessageDTO struct {
	Id                   int64     `json:"id"`
	SenderId             uuid.UUID `json:"senderId"`
	ReceiverId           uuid.UUID `json:"receiverId"`
	SenderEmailAddress   string    `json:"senderEmailAddress"`
	ReceiverEmailAddress string    `json:"receiverEmailAddress"`
	Subject              string    `json:"subject"`
	Body                 string    `json:"body"`
	IsHtml               bool      `json:"isHtml"`
	SentDate             time.Time `json:"sentDate"`
}
