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
	LastModifiedOn time.Time  `gorm:"not null;type:timestamp with time zone;default:now()"`
	LastModifiedBy uuid.UUID  `gorm:"type:uuid;not null"`
	DeletedOn      *time.Time `gorm:"type:timestamp with time zone"`
	DeletedBy      *uuid.UUID `gorm:"type:uuid"`
}

// BeforeSave hook - BaseEntity kullanan tüm modeller için otomatik timestamp güncelleme
func (b *BaseEntity) BeforeSave(tx *gorm.DB) error {
	b.LastModifiedOn = time.Now()
	return nil
}

// BeforeCreate hook - BaseEntity kullanan tüm modeller için otomatik CreatedOn set etme
func (b *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreatedOn = now
	b.LastModifiedOn = now
	return nil
}

// Soft delete kontrolü
func (b *BaseEntity) BeforeDelete(tx *gorm.DB) error {
	return tx.Model(b).Update("DeletedOn", time.Now()).Error
}
