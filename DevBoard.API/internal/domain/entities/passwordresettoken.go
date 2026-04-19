// internal/domain/entities/PasswordResetToken.go
package entities

import (
	"time"

	"github.com/google/uuid"
)

type PasswordResetToken struct {
	Id        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserId    uuid.UUID  `gorm:"type:uuid;not null"`
	Token     string     `gorm:"type:varchar(500);not null;unique"`
	ExpiresAt time.Time  `gorm:"type:timestamp with time zone;not null"`
	UsedAt    *time.Time `gorm:"type:timestamp with time zone"`
	CreatedOn time.Time  `gorm:"type:timestamp with time zone;not null;default:now()"`

	User *User `gorm:"foreignKey:UserId"`
}

func (PasswordResetToken) TableName() string { return "PasswordResetTokens" }

// IsValid checks if the token is not used and not expired
func (p *PasswordResetToken) IsValid() bool {
	return p.UsedAt == nil && time.Now().Before(p.ExpiresAt)
}
