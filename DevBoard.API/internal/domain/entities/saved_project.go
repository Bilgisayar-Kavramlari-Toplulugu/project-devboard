package entities

import (
	"github.com/google/uuid"
)

type SavedProject struct {
	Id        int64     `gorm:"type:bigint;primaryKey"`
	UserId    uuid.UUID `gorm:"type:uuid;not null"`
	ProjectId int64     `gorm:"type:bigint;not null"`
	BaseEntity

	User    User    `gorm:"foreignKey:UserId;references:Id"`
	Project Project `gorm:"foreignKey:ProjectId;references:Id"`
}

func (SavedProject) TableName() string { return "SavedProjects" }
