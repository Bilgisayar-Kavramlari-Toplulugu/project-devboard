package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User - Kullanıcılar
type User struct {
	// Props
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email             string    `gorm:"type:text;uniqueIndex;not null"`
	Password          string    `gorm:"type:text"`
	Firstname         string    `gorm:"type:text;not null"`
	Lastname          string    `gorm:"type:varchar(500)"`
	IsEmailValidated  bool      `gorm:"type:boolean;not null;default:false"`
	PhoneNumber       *string   `gorm:"type:varchar(500)"`
	ProfilePictureUrl *string   `gorm:"type:text"`

	// Session properties
	RefreshTokenHash *string    `gorm:"type:text"`
	RefreshTokenExp  *time.Time
	DeviceInfo       []byte     `gorm:"type:jsonb"`
	IPAddress        *string    `gorm:"type:varchar(500)"`
	UserAgent        *string    `gorm:"type:text"`

	BaseEntity

	// Relations
	UserRoles []UserRole `gorm:"foreignKey:UserID"`
}

func (User) TableName() string { return "core.Users" }

// BeforeCreate hooks - UUID otomatik oluşturma
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
