package entities

import (
	"github.com/google/uuid"
)

type SavedFilter struct {
	Id         int64     `gorm:"type:bigint;primaryKey"`
	Name       string    `gorm:"type:varchar(500);not null"`
	FilterData string    `gorm:"type:jsonb;not null"`
	OwnerId    uuid.UUID `gorm:"type:uuid;not null"`
	BaseEntity

	Owner User `gorm:"foreignKey:OwnerId;references:Id"`
}

func (SavedFilter) TableName() string { return "SavedFilters" }
