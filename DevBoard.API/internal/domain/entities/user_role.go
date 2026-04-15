package entities

import (
	"github.com/google/uuid"
)

type UserRole struct {
	UserId uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleId int       `gorm:"type:integer;primaryKey"`
	BaseEntity

	User User `gorm:"foreignKey:UserId;references:Id"`
	Role Role `gorm:"foreignKey:RoleId;references:Id"`
}

func (UserRole) TableName() string { return "UserRoles" }
