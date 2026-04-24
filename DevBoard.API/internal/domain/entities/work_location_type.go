package entities

type WorkLocationType struct {
	Id   int     `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null;uniqueIndex"`
	BaseEntity

	UserWorkLocationTypes []UserWorkLocationType `gorm:"foreignKey:WorkLocationTypeId"`
}

func (WorkLocationType) TableName() string { return "WorkLocationTypes" }
