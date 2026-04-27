package config

import (
	"log"
	"project-devboard/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var SystemUserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

func SeedData(db *gorm.DB) {
	seedJobTypes(db)
	seedCountries(db)
	seedCities(db)
}

func isSeeded(db *gorm.DB, model interface{}) bool {
    var count int64
    db.Model(model).Count(&count)
    return count > 0
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
func seedCountries(db *gorm.DB) {
    if isSeeded(db, &entities.Country{}) {
        return
    }
    log.Println("🌱 Seeding countries...")
    data := getCountrySeedData()
    if err := db.CreateInBatches(data, 50).Error; err != nil {
        log.Printf("❌ Country seed error: %v\n", err)
    }
    log.Printf("✅ %d countries seeded\n", len(data))
}

func seedCities(db *gorm.DB) {
    if isSeeded(db, &entities.City{}) {
        return
    }
    log.Println("🌱 Seeding cities...")
    data := getCitySeedData()
    if err := db.CreateInBatches(data, 100).Error; err != nil {
        log.Printf("❌ City seed error: %v\n", err)
    }
    log.Printf("✅ %d cities seeded\n", len(data))
}