package routers

import (
	"personal/app/controllers"
	"personal/app/middlewares"
	"personal/app/repositories"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

var (
	portfolioRepository repositories.PortfolioRepository = repositories.NewPortfolioRepository(db)

	portfolioService services.PortfolioService = services.NewPortfolioService(portfolioRepository)

	portfolioController controllers.PortfolioController = controllers.NewPortfolioController(portfolioService)
)

func PortfolioRoute(route *gin.Engine) {

	portfolioRoutes := route.Group("api/v1/portfolio")
	{
		portfolioRoutes.GET("/", portfolioController.All)
		portfolioRoutes.GET("/:id", portfolioController.FindByID)
		portfolioRoutes.Use(middlewares.JwtAuthMiddleware())
		portfolioRoutes.POST("/", portfolioController.Insert)
		portfolioRoutes.PUT("/:id", portfolioController.Update)
		portfolioRoutes.DELETE("/:id", portfolioController.Delete)
		portfolioRoutes.GET("/deleted", portfolioController.Deleted)
		portfolioRoutes.PUT("/restore/:id", portfolioController.Restore)
		portfolioRoutes.DELETE("/delete/:id", portfolioController.DeletePermanent)
	}
}
