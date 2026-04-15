package entities

type Skill struct {
	Id          int    `gorm:"type:integer;primaryKey"`
	Name        string `gorm:"type:varchar(500);not null"`
	SkillTypeId int    `gorm:"type:integer;not null"`
	BaseEntity

	SkillType               SkillType                 `gorm:"foreignKey:SkillTypeId;references:Id"`
	UserSkills              []UserSkill               `gorm:"foreignKey:SkillId"`
	ProjectEndorsementables []ProjectEndorsementable  `gorm:"foreignKey:SkillId"`
}

func (Skill) TableName() string { return "Skills" }
