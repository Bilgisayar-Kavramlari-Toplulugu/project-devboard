package entities

type ProjectRole struct {
	Id   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null"`
	BaseEntity

	Projects           []Project           `gorm:"foreignKey:OwnerProjectRole"`
	ProjectDevelopers  []ProjectDeveloper  `gorm:"foreignKey:ProjectRoleId"`
}

func (ProjectRole) TableName() string { return "ProjectRoles" }
