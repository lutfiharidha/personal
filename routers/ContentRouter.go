package routers

import (
	"personal/app/controllers"
	"personal/app/middlewares"
	"personal/app/repositories"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

var (
	contentRepository repositories.ContentRepository = repositories.NewContentRepository(db)

	contentService services.ContentService = services.NewContentService(contentRepository)

	contentController controllers.ContentController = controllers.NewContentController(contentService)
)

func ContentRoute(route *gin.Engine) {

	contentRoutes := route.Group("api/v1/content")
	{
		contentRoutes.Use(middlewares.JwtAuthMiddleware())
		contentRoutes.GET("/", contentController.All)
		contentRoutes.POST("/", contentController.Insert)
		contentRoutes.GET("/:id", contentController.FindByID)
		contentRoutes.PUT("/:id", contentController.Update)
		contentRoutes.DELETE("/:id", contentController.Delete)
	}
}
