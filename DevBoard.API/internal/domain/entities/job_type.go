package entities

type JobType struct {
	Id   int    `gorm:"type:integer;primaryKey"`
	Name string `gorm:"type:varchar(500);not null"`
	BaseEntity

	UserJobTypes []UserJobType `gorm:"foreignKey:JobTypeId"`
}

func (JobType) TableName() string { return "JobTypes" }
