package entities

type ProjectEndorsementable struct {
	Id        int64  `gorm:"type:bigint;primaryKey"`
	ProjectId *int64 `gorm:"type:bigint"`
	SkillId   *int64 `gorm:"type:bigint"`
	BaseEntity

	Project             *Project              `gorm:"foreignKey:ProjectId;references:Id"`
	ProjectEndorsements []ProjectEndorsement  `gorm:"foreignKey:EndorsementableId"`
}

func (ProjectEndorsementable) TableName() string { return "ProjectEndorsementables" }
