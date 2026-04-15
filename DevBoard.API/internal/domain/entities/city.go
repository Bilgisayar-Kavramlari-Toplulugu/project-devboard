package entities

type City struct {
	Id           int     `gorm:"type:integer;primaryKey"`
	Name         string  `gorm:"type:varchar(500);not null"`
	Code         *string `gorm:"type:varchar(500)"`
	DisplayOrder *int    `gorm:"type:integer"`
	CountryId    *int    `gorm:"type:integer"`
	BaseEntity

	Country     *Country      `gorm:"foreignKey:CountryId;references:Id"`
	Users       []User        `gorm:"foreignKey:CityId"`
	Experiences []Experience  `gorm:"foreignKey:CityId"`
	Educations  []Education   `gorm:"foreignKey:CityId"`
}

func (City) TableName() string { return "Cities" }
