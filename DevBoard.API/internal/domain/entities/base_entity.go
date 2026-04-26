package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseEntity - Tüm tablolar için ortak alanlar
type BaseEntity struct {
	IsActive       bool       `gorm:"type:boolean;not null;default:true"`
	CreatedOn      time.Time  `gorm:"not null;type:timestamp with time zone;default:now()"`
	CreatedBy      uuid.UUID  `gorm:"type:uuid;not null"`
	LastModifiedAt time.Time  `gorm:"not null;type:timestamp with time zone;default:now()"`
	LastModifiedBy uuid.UUID  `gorm:"type:uuid;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	DeletedBy      *uuid.UUID `gorm:"type:uuid"`
}

// BeforeSave hook - BaseEntity kullanan tüm modeller için otomatik timestamp güncelleme
// bu herzaman çalışıyor !!!!!!!!!!!!!!!!!
func (b *BaseEntity) BeforeSave(tx *gorm.DB) error {
	b.LastModifiedAt = time.Now()
	return nil
}

// BeforeCreate hook - BaseEntity kullanan tüm modeller için otomatik CreatedOn set etme
// bu ikisi tek fonksiyona birleştirilebilir
func (b *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreatedOn = now
	b.LastModifiedAt = now
	return nil
}
