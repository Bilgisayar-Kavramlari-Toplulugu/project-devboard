package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PublicEndorsementDTO struct {
	Id                  int64     `json:"id"`
	UserSkillId         int64     `json:"userSkillId"`
	SenderUserId        uuid.UUID `json:"senderUserId"`
	SenderUserFirstName string    `json:"senderUserFirstName"`
	SenderUserLastName  string    `json:"senderUserLastName"`
	SenderEmailAddress  string    `json:"senderEmailAddress"`
	CreatedDate         time.Time `json:"createdDate"`
}
