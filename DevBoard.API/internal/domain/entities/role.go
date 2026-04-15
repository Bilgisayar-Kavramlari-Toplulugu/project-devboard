package entities

type Role struct {
	Id   int   `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null"`
	BaseEntity

	UserRoles []UserRole `gorm:"foreignKey:RoleId"`
}

func (Role) TableName() string { return "Roles" }
