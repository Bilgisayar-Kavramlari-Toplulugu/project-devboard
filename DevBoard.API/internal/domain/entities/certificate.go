package entities

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	Id         int64      `gorm:"type:bigint;primaryKey"`
	UserId     uuid.UUID  `gorm:"type:uuid;not null"`
	Name       string     `gorm:"type:varchar(500);not null"`
	Degree     *string    `gorm:"type:varchar(500)"`
	IssueDate  time.Time  `gorm:"type:timestamp;not null"`
	ExpireDate *time.Time `gorm:"type:timestamp"`
	Url        *string    `gorm:"type:varchar(500)"`
	BaseEntity

	User User `gorm:"foreignKey:UserId;references:Id"`
}

func (Certificate) TableName() string { return "Certificates" }
