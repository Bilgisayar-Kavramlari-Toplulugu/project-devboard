package dtos

import "github.com/google/uuid"

type ProjectEndorsementDTO struct {
	Id                int64     `json:"id"`
	EndorsementableId int64     `json:"endorsementableId"`
	SenderId          uuid.UUID `json:"senderId"`
}
