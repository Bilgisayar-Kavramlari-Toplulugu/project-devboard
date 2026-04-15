package db_plugins

import (
	"time"

	"gorm.io/gorm"
)

// Ayrı bir dosyada: plugins/timestamp_plugin.go
type TimestampPlugin struct{}

func (p *TimestampPlugin) Name() string {
	return "TimestampPlugin"
}

func (p *TimestampPlugin) Initialize(db *gorm.DB) error {
	// BeforeCreate callback
	db.Callback().Create().Before("gorm:create").Register("timestamp:before_create", func(db *gorm.DB) {
		if db.Statement.Schema != nil {
			now := time.Now()
			if field := db.Statement.Schema.LookUpField("CreatedOn"); field != nil {
				db.Statement.SetColumn("CreatedOn", now)
			}
			if field := db.Statement.Schema.LookUpField("LastModifiedOn"); field != nil {
				db.Statement.SetColumn("LastModifiedOn", now)
			}
		}
	})

	// BeforeUpdate callback
	db.Callback().Update().Before("gorm:update").Register("timestamp:before_update", func(db *gorm.DB) {
		if db.Statement.Schema != nil {
			if field := db.Statement.Schema.LookUpField("LastModifiedOn"); field != nil {
				db.Statement.SetColumn("LastModifiedOn", time.Now())
			}
		}
	})

	return nil
}
