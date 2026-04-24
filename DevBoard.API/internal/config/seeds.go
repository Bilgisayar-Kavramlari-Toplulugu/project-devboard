package config

import (
	"project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var SystemUserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

func SeedData(db *gorm.DB) {
	seedJobTypes(db)
	seedWorkLocationTypes(db)
}

func seedJobTypes(db *gorm.DB) {
	jobTypes := []entities.JobType{
		{Name: "Full-time", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "Part-time", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "Contract", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "Temporary", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "Internship", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "Other", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&jobTypes)
}

func seedWorkLocationTypes(db *gorm.DB) {
	workLocationTypes := []entities.WorkLocationType{
		{Name: "Remote", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "Hybrid", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
		{Name: "OnSite", BaseEntity: entities.BaseEntity{CreatedBy: SystemUserID, LastModifiedBy: SystemUserID}},
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&workLocationTypes)
}
