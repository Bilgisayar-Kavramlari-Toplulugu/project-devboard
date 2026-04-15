package entities

type WorkLocationType struct {
	Id   int     `gorm:"type:integer;primaryKey"`
	Name *string `gorm:"type:varchar(500)"`
	BaseEntity

	UserWorkLocationTypes []UserWorkLocationType `gorm:"foreignKey:WorkLocationTypeId"`
}

func (WorkLocationType) TableName() string { return "WorkLocaitonTypes" }
