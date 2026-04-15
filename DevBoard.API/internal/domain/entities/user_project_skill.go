package entities

type UserProjectSkill struct {
	Id          int64 `gorm:"type:bigint;primaryKey"`
	ProjectId   int64 `gorm:"type:bigint;not null"`
	UserSkillId int64 `gorm:"type:bigint;not null"`
	BaseEntity

	Project   Project   `gorm:"foreignKey:ProjectId;references:Id"`
	UserSkill UserSkill `gorm:"foreignKey:UserSkillId;references:Id"`
}

func (UserProjectSkill) TableName() string { return "UserProjectSkills" }
