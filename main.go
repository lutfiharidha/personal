package main

import (
	"personal/app/handlers"
	"personal/configs"

	"gorm.io/gorm"
)

var db *gorm.DB = configs.SetupDatabaseConnection()

func main() {
	defer configs.CloseDatabaseConnection(db)
	handlers.Routing()
}
