package entities

type MessageTemplateType struct {
	Id   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null"`
	BaseEntity

	MessageTemplates []MessageTemplate `gorm:"foreignKey:TypeId"`
	Messages         []Message         `gorm:"foreignKey:TypeId"`
}

func (MessageTemplateType) TableName() string { return "MessageTemplateTypes" }
