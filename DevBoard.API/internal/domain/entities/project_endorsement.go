package entities

import (
	"github.com/google/uuid"
)

type ProjectEndorsement struct {
	Id               int64     `gorm:"type:bigint;primaryKey"`
	EndorsementableId int64    `gorm:"type:bigint;not null"`
	SenderId         uuid.UUID `gorm:"type:uuid;not null"`
	BaseEntity

	Endorsementable ProjectEndorsementable `gorm:"foreignKey:EndorsementableId;references:Id"`
	Sender          User                   `gorm:"foreignKey:SenderId;references:Id"`
}

func (ProjectEndorsement) TableName() string { return "ProjectEndorsements" }
