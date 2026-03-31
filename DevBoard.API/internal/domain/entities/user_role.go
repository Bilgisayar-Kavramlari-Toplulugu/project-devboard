package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole - Kullanıcı rolleri (many-to-many)
type UserRole struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	RoleID uuid.UUID `gorm:"type:uuid;not null"`

	BaseEntity

	// Relations
	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}

func (UserRole) TableName() string { return "core.UserRoles" }

func (ur *UserRole) BeforeCreate(tx *gorm.DB) error {
	if ur.ID == uuid.Nil {
		ur.ID = uuid.New()
	}
	return nil
}
