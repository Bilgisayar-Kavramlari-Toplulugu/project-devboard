package config

import (
	"log"

	domain "project-devboard/internal/domain/entities"

	"gorm.io/gorm"
)

// RunMigrations - Tüm domain modellerini migrate eder
func RunMigrations(db *gorm.DB) {
	// Önce schema'ları oluştur
	createSchemas(db)

	// Seviye 0: Bağımsız tablolar (hiçbir FK referansı yok)
	if err := db.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.JobType{},
		&domain.SkillType{},
		&domain.WorkLocationType{},
	); err != nil {
		log.Fatal("Migration Failed (Level 0 - Independent tables):", err)
	}

	// Seviye 1: Sadece Seviye 0'a bağımlı tablolar
	if err := db.AutoMigrate(
		&domain.UserRole{}, // -> User, Role
		&domain.UserJobType{},
		&domain.PasswordResetToken{},
	); err != nil {
		log.Fatal("Migration Failed (Level 1 - User dependent tables):", err)
	}

	// Seviye 2: Seviye 0 ve 1'e bağımlı tablolar

	// Seviye 3: Seviye 2'ye bağımlı tablolar

	// Seed data
	SeedData(db)

	log.Println("Migrations completed successfully")
}

// createSchemas - Gerekli schema'ları oluşturur
func createSchemas(db *gorm.DB) {
	schemas := []DBSchema{Core, App}

	for _, schema := range schemas {
		sql := "CREATE SCHEMA IF NOT EXISTS " + schema.String()
		if err := db.Exec(sql).Error; err != nil {
			log.Printf("Warning: Could not create schema %s: %v", schema, err)
		}
	}

	log.Println("Schemas created/verified successfully")
}
