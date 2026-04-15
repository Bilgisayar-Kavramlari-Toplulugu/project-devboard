package entities

import (
	"github.com/google/uuid"
)

type UserJobType struct {
	UserId    uuid.UUID `gorm:"type:uuid;primaryKey"`
	JobTypeId int       `gorm:"type:integer;primaryKey"`
	BaseEntity

	User    User    `gorm:"foreignKey:UserId;references:Id"`
	JobType JobType `gorm:"foreignKey:JobTypeId;references:Id"`
}

func (UserJobType) TableName() string { return "UserJobTypes" }
