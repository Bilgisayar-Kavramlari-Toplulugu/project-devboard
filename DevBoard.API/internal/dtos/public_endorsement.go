package dtos

import "github.com/google/uuid"

type PublicEndorsementDTO struct {
	Id           int64     `json:"id"`
	UserSkillId  int64     `json:"userSkillId"`
	SenderUserId uuid.UUID `json:"senderUserId"`
}
