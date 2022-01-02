package main

import (
	"os"
	"personal/configs"
	"personal/routers"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = configs.SetupDatabaseConnection()

func main() {
	defer configs.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}))
	r.Static("/image/client", "/root/Go/src/personal/public/img/client")
	r.Static("/image/portfolio", "/root/Go/src/personal/public/img/portfolio")
	routers.ClientRoute(r)    //Added all client routes
	routers.AuthRoute(r)      //Added all auth routes
	routers.ProfileRoute(r)   //Added all profile routes
	routers.ContactRoute(r)   //Added all contact routes
	routers.PortfolioRoute(r) //Added all portfolio routes
	routers.SkillRoute(r)     //Added all skill routes
	routers.ContentRoute(r)   //Added all skill routes

	r.Run(":" + os.Getenv("APP_PORT"))
}
