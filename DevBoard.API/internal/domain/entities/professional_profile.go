package entities

import (
	"github.com/google/uuid"
)

type ProfessionalProfile struct {
	Id         int       `gorm:"type:integer;primaryKey"`
	UserId     uuid.UUID `gorm:"type:uuid;not null"`
	PlatformId int       `gorm:"type:integer;not null"`
	Url        string    `gorm:"type:varchar(500);not null"`
	BaseEntity

	User     User                 `gorm:"foreignKey:UserId;references:Id"`
	Platform ProfessionalPlatform `gorm:"foreignKey:PlatformId;references:Id"`
}

func (ProfessionalProfile) TableName() string { return "ProfessionalProfiles" }
