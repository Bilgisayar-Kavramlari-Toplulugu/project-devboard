package dtos

import "github.com/google/uuid"

type ProjectEndorsementDTO struct {
	Id                  int64     `json:"id"`
	EndorsementableId   int64     `json:"endorsementableId"`
	EndorsementableName string    `json:"endorsementableName"`
	SenderId            uuid.UUID `json:"senderId"`
}
