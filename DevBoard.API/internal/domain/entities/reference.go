package entities

import (
	"github.com/google/uuid"
)

type Reference struct {
	Id           int64      `gorm:"type:bigint;primaryKey"`
	UserId       *uuid.UUID `gorm:"type:uuid"`
	Firstname    string     `gorm:"type:varchar(500);not null"`
	Lastname     string     `gorm:"type:varchar(500);not null"`
	PhoneNumber  *string    `gorm:"type:varchar(500)"`
	EmailAddress *string    `gorm:"type:varchar(500)"`
	BaseEntity

	User *User `gorm:"foreignKey:UserId;references:Id"`
}

func (Reference) TableName() string { return "References" }
