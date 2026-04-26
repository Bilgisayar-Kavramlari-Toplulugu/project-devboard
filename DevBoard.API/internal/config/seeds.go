package config

import (
	"project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var SystemUserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

func SeedData(db *gorm.DB) {
	seedJobTypes(db)
}

func seedJobTypes(db *gorm.DB) {
	jobTypes := []entities.JobType{
		{Id: 1, Name: "Full-time", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Id: 2, Name: "Part-time", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Id: 3, Name: "Contract", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Id: 4, Name: "Temporary", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Id: 5, Name: "Internship", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Id: 6, Name: "Other", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
	}

	for _, jt := range jobTypes {
		var count int64
		db.Model(&entities.JobType{}).Where("id = ?", jt.Id).Count(&count)
		if count == 0 {
			db.Create(&jt)
		}
	}
}
