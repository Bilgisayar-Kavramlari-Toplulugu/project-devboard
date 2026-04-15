package entities

type ProjectType struct {
	Id   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null"`
	BaseEntity

	Projects []Project `gorm:"foreignKey:TypeId"`
}

func (ProjectType) TableName() string { return "ProjectTypes" }
