package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PasswordResetToken - Şifre sıfırlama tokenları
type PasswordResetToken struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	TokenHash string     `gorm:"type:text;not null;uniqueIndex"`
	ExpiresAt time.Time  `gorm:"not null;type:timestamp with time zone"`
	UsedAt    *time.Time `gorm:"type:timestamp with time zone"`
	CreatedAt time.Time  `gorm:"not null;type:timestamp with time zone"`

	// Relations
	User User `gorm:"foreignKey:UserID"`
}

func (PasswordResetToken) TableName() string { return "core.PasswordResetTokens" }

func (p *PasswordResetToken) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	return nil
}

// IsValid - Token geçerli mi? (süresi dolmamış ve kullanılmamış)
func (p *PasswordResetToken) IsValid() bool {
	return p.UsedAt == nil && p.ExpiresAt.After(time.Now())
}
