package entities

import (
	"github.com/google/uuid"
)

type UserSkill struct {
	Id      int64     `gorm:"type:bigint;primaryKey"`
	UserId  uuid.UUID `gorm:"type:uuid;not null"`
	SkillId int       `gorm:"type:integer;not null"`
	BaseEntity

	User                  User                 `gorm:"foreignKey:UserId;references:Id"`
	Skill                 Skill                `gorm:"foreignKey:SkillId;references:Id"`
	PublicEndorsements    []PublicEndorsement  `gorm:"foreignKey:UserSkillId"`
	UserProjectSkills     []UserProjectSkill   `gorm:"foreignKey:UserSkillId"`
}

func (UserSkill) TableName() string { return "UserSkills" }
