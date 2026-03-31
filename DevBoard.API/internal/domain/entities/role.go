package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role - Roller
type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string    `gorm:"type:text;not null"`

	BaseEntity

	// Relations
	UserRoles []UserRole `gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string { return "core.Roles" }

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
