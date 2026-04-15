package entities

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id               int64      `gorm:"type:bigint;primaryKey"`
	TypeId           int        `gorm:"type:integer;not null"`
	UserId           *uuid.UUID `gorm:"type:uuid"`
	OwnerProjectRole *int       `gorm:"type:integer"`
	Name             string     `gorm:"type:varchar(500);not null"`
	ShortDescription string     `gorm:"type:varchar(500);not null"`
	Description      string     `gorm:"type:varchar(500);not null"`
	StartDate        time.Time  `gorm:"type:timestamp;not null"`
	EndDate          *time.Time `gorm:"type:timestamp"`
	Url              *string    `gorm:"type:varchar(500)"`
	BaseEntity

	Type                    ProjectType               `gorm:"foreignKey:TypeId;references:Id"`
	User                    *User                     `gorm:"foreignKey:UserId;references:Id"`
	OwnerRole               *ProjectRole              `gorm:"foreignKey:OwnerProjectRole;references:Id"`
	UserProjectSkills       []UserProjectSkill        `gorm:"foreignKey:ProjectId"`
	ProjectDevelopers       []ProjectDeveloper        `gorm:"foreignKey:ProjectId"`
	ProjectEndorsementables []ProjectEndorsementable  `gorm:"foreignKey:ProjectId"`
	SavedProjects           []SavedProject            `gorm:"foreignKey:ProjectId"`
}

func (Project) TableName() string { return "Projects" }
