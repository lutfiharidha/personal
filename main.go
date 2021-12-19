package main

import (
	"os"
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

	routers.ClientRoute(r)  //Added all client routes
	routers.AuthRoute(r)    //Added all auth routes
	routers.ProfileRoute(r) //Added all profile routes

	r.Run(":" + os.Getenv("APP_PORT"))
}
