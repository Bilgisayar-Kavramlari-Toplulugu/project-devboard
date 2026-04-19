package entities

import (
	"github.com/google/uuid"
)

type Message struct {
	Id         int64     `gorm:"type:bigint;primaryKey"`
	SenderId   uuid.UUID `gorm:"type:uuid;not null"`
	ReceiverId uuid.UUID `gorm:"type:uuid;not null"`
	Subject    string    `gorm:"type:varchar(500);not null"`
	Body       string    `gorm:"type:varchar(500);not null"`
	IsHtml     bool      `gorm:"type:boolean;not null"`
	TypeId     *int      `gorm:"type:integer"`
	BaseEntity

	Sender   User                 `gorm:"foreignKey:SenderId;references:Id"`
	Receiver User                 `gorm:"foreignKey:ReceiverId;references:Id"`
	Type     *MessageTemplateType `gorm:"foreignKey:TypeId;references:Id"`
}

func (Message) TableName() string { return "Messages" }
