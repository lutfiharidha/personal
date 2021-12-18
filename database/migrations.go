package database

import (
	"personal/app/configs"
	"personal/app/models"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.SetupDatabaseConnection()
)

func InitialMigration() {
	db.Migrator().DropTable(&models.Client{})
	db.Migrator().CreateTable(&models.Client{})
}
