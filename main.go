package main

import (
	"personal/app/configs"
	"personal/app/handlers"
	"personal/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = configs.SetupDatabaseConnection()

func main() {
	defer configs.CloseDatabaseConnection(db)
	handlers.Command()
	r := gin.New()
	r.Use(configs.CORSMiddleware())

	routers.ClientRoute(r) //Added all product routes

	r.Run()
}
