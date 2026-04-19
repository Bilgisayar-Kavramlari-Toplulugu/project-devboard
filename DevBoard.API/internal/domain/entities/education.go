package entities

import (
	"github.com/google/uuid"
)

type Education struct {
	Id     int64      `gorm:"type:bigint;primaryKey"`
	UserId *uuid.UUID `gorm:"type:uuid"`
	Name   string     `gorm:"type:varchar(500);not null"`
	CityId int        `gorm:"type:integer;not null"`
	GDPR   *float64   `gorm:"type:double precision"`
	BaseEntity

	User *User `gorm:"foreignKey:UserId;references:Id"`
	City City  `gorm:"foreignKey:CityId;references:Id"`
}

func (Education) TableName() string { return "Educations" }
