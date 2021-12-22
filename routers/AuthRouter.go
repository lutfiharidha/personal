package routers

import (
	"personal/app/controllers"
	"personal/app/repositories"
	"personal/app/services"
	"personal/configs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.SetupDatabaseConnection()

	authRepository repositories.AuthRepository = repositories.NewAuthRepository(db)

	authService services.AuthService = services.NewAuthService(authRepository)

	authController controllers.AuthController = controllers.NewAuthController(authService)
)

func AuthRoute(route *gin.Engine) {

	authRoutes := route.Group("api/v1/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}
}
