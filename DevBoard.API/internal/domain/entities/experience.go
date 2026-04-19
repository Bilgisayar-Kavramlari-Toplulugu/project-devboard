package entities

import (
	"time"

	"github.com/google/uuid"
)

type Experience struct {
	Id          int64      `gorm:"type:bigint;primaryKey"`
	UserId      *uuid.UUID `gorm:"type:uuid"`
	Name        string     `gorm:"type:varchar(500);not null"`
	CityId      int        `gorm:"type:integer;not null"`
	Startdate   time.Time  `gorm:"type:timestamp;not null"`
	Enddate     *time.Time `gorm:"type:timestamp"`
	Position    *string    `gorm:"type:varchar(500)"`
	Information *int64     `gorm:"type:bigint"`
	BaseEntity

	User *User `gorm:"foreignKey:UserId;references:Id"`
	City City  `gorm:"foreignKey:CityId;references:Id"`
}

func (Experience) TableName() string { return "Experiences" }
