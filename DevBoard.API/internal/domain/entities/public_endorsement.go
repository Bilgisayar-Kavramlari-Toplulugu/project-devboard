package entities

import (
	"github.com/google/uuid"
)

type PublicEndorsement struct {
	Id           int64     `gorm:"type:bigint;primaryKey"`
	UserSkillId  int64     `gorm:"type:bigint;not null"`
	SenderUserId uuid.UUID `gorm:"type:uuid;not null"`
	BaseEntity

	UserSkill  UserSkill `gorm:"foreignKey:UserSkillId;references:Id"`
	SenderUser User      `gorm:"foreignKey:SenderUserId;references:Id"`
}

func (PublicEndorsement) TableName() string { return "PublicEndorsements" }
