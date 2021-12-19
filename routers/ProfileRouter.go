package routers

import (
	"personal/app/controllers"
	"personal/app/middlewares"
	"personal/app/repositories"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

var (
	profileRepository repositories.ProfileRepository = repositories.NewProfileRepository(db)

	profileService services.ProfileService = services.NewProfileService(profileRepository)

	profileController controllers.ProfileController = controllers.NewProfileController(profileService)
)

func ProfileRoute(route *gin.Engine) {

	profileRoutes := route.Group("api/v1/profile")
	{
		profileRoutes.Use(middlewares.JwtAuthMiddleware())

		profileRoutes.GET("/", profileController.Profile)
		profileRoutes.POST("/logout", profileController.Logout)
	}
}
