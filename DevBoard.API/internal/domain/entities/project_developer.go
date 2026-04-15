package entities

import (
	"github.com/google/uuid"
)

type ProjectDeveloper struct {
	Id            int64     `gorm:"type:bigint;primaryKey"`
	ProjectId     int64     `gorm:"type:bigint;not null"`
	DeveloperId   uuid.UUID `gorm:"type:uuid;not null"`
	ProjectRoleId int       `gorm:"type:integer;not null"`
	BaseEntity

	Project     Project     `gorm:"foreignKey:ProjectId;references:Id"`
	Developer   User        `gorm:"foreignKey:DeveloperId;references:Id"`
	ProjectRole ProjectRole `gorm:"foreignKey:ProjectRoleId;references:Id"`
}

func (ProjectDeveloper) TableName() string { return "ProjectDevelopers" }
