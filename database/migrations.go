package database

import (
	"personal/app/models"
	"personal/configs"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.SetupDatabaseConnection()
)

func InitialMigration() {
	db.Migrator().DropTable(&models.Client{})
	db.Migrator().CreateTable(&models.Client{})

	db.Migrator().DropTable(&models.Auth{})
	db.Migrator().CreateTable(&models.Auth{})

	db.Migrator().DropTable(&models.Contact{})
	db.Migrator().CreateTable(&models.Contact{})

	db.Migrator().DropTable(&models.Portfolio{})
	db.Migrator().CreateTable(&models.Portfolio{})

	db.Migrator().DropTable(&models.Skill{})
	db.Migrator().CreateTable(&models.Skill{})

	db.Migrator().DropTable(&models.Content{})
	db.Migrator().CreateTable(&models.Content{})
}
