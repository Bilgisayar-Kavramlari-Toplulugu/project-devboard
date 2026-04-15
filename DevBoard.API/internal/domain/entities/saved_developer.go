package entities

import (
	"github.com/google/uuid"
)

type SavedDeveloper struct {
	Id          int64     `gorm:"type:bigint;primaryKey"`
	UserId      uuid.UUID `gorm:"type:uuid;not null"`
	DeveloperId uuid.UUID `gorm:"type:uuid;not null"`
	BaseEntity

	User      User `gorm:"foreignKey:UserId;references:Id"`
	Developer User `gorm:"foreignKey:DeveloperId;references:Id"`
}

func (SavedDeveloper) TableName() string { return "SavedDevelopers" }
