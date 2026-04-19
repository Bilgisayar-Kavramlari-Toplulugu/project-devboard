package entities

type MessageTemplate struct {
	Id     int64 `gorm:"type:bigint;primaryKey"`
	TypeId int   `gorm:"type:integer;not null"`
	Subject int64 `gorm:"type:bigint;not null"`
	Body    int64 `gorm:"type:bigint;not null"`
	IsHtml bool  `gorm:"type:boolean;not null"`
	BaseEntity

	Type MessageTemplateType `gorm:"foreignKey:TypeId;references:Id"`
}

func (MessageTemplate) TableName() string { return "MessageTemplates" }
