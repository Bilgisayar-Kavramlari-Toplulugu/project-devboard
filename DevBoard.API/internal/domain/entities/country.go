package entities

type Country struct {
	Id          int    `gorm:"type:integer;primaryKey"`
	Name        string `gorm:"type:varchar(500);not null"`
	FlagCode    string `gorm:"type:varchar(500);not null"`
	ShortCode   string `gorm:"type:varchar(500);not null"`
	PhonePrefix string `gorm:"type:varchar(500);not null"`
	BaseEntity

	Cities []City `gorm:"foreignKey:CountryId"`
}

func (Country) TableName() string { return "Countries" }
