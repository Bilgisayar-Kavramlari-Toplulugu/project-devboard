package entities

type SkillType struct {
	Id   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null"`
	BaseEntity

	Skills []Skill `gorm:"foreignKey:SkillTypeId"`
}

func (SkillType) TableName() string { return "SkillTypes" }
