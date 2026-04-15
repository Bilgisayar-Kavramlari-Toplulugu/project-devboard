package entities

import (
	"github.com/google/uuid"
)

type UserWorkLocationType struct {
	UserId             uuid.UUID `gorm:"type:uuid;primaryKey"`
	WorkLocationTypeId int       `gorm:"type:integer;primaryKey"`
	BaseEntity

	User             User             `gorm:"foreignKey:UserId;references:Id"`
	WorkLocationType WorkLocationType `gorm:"foreignKey:WorkLocationTypeId;references:Id"`
}

func (UserWorkLocationType) TableName() string { return "UserWorkLocationTypes" }
