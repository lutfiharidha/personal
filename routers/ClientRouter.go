package routers

import (
	"personal/app/controllers"
	"personal/app/middlewares"
	"personal/app/repositories"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

var (
	clientRepository repositories.ClientRepository = repositories.NewClientRepository(db)

	clientService services.ClientService = services.NewClientService(clientRepository)

	clientController controllers.ClientController = controllers.NewClientController(clientService)
)

func ClientRoute(route *gin.Engine) {

	clientRoutes := route.Group("api/v1/client")
	{
		clientRoutes.Use(middlewares.JwtAuthMiddleware())
		clientRoutes.GET("/", clientController.All)
		clientRoutes.POST("/", clientController.Insert)
		clientRoutes.GET("/:id", clientController.FindByID)
		clientRoutes.PUT("/:id", clientController.Update)
		clientRoutes.DELETE("/:id", clientController.Delete)
		clientRoutes.GET("/deleted", clientController.Deleted)
		clientRoutes.PUT("/restore/:id", clientController.Restore)
		clientRoutes.DELETE("/delete/:id", clientController.DeletePermanent)
	}
}
