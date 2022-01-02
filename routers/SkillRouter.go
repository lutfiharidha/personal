package routers

import (
	"personal/app/controllers"
	"personal/app/middlewares"
	"personal/app/repositories"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

var (
	skillRepository repositories.SkillRepository = repositories.NewSkillRepository(db)

	skillService services.SkillService = services.NewSkillService(skillRepository)

	skillController controllers.SkillController = controllers.NewSkillController(skillService)
)

func SkillRoute(route *gin.Engine) {

	skillRoutes := route.Group("api/v1/skill")
	{
		skillRoutes.GET("/", skillController.All)
		skillRoutes.Use(middlewares.JwtAuthMiddleware())
		skillRoutes.POST("/", skillController.Insert)
		skillRoutes.DELETE("/:id", skillController.Delete)
	}
}
