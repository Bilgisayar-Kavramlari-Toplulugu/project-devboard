package entities

type ProfessionalPlatform struct {
	Id   int     `gorm:"type:integer;primaryKey"`
	Name *string `gorm:"type:varchar(500)"`
	BaseEntity

	ProfessionalProfiles []ProfessionalProfile `gorm:"foreignKey:PlatformId"`
}

func (ProfessionalPlatform) TableName() string { return "ProfessionalPlatforms" }
