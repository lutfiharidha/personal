package handlers

import (
	"os"
	"personal/configs"
	"personal/routers"

	"github.com/gin-gonic/gin"
)

func Routing() {
	r := gin.New()
	r.Use(configs.CORSMiddleware())

	routers.ClientRoute(r)    //Added all client routes
	routers.AuthRoute(r)      //Added all auth routes
	routers.ProfileRoute(r)   //Added all profile routes
	routers.ContactRoute(r)   //Added all contact routes
	routers.PortfolioRoute(r) //Added all portfolio routes
	routers.SkillRoute(r)     //Added all skill routes
	routers.ContentRoute(r)   //Added all skill routes

	r.Run(":" + os.Getenv("APP_PORT"))
}
