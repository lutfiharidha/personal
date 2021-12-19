package controllers

import (
	"net/http"
	"personal/app/dtos"
	"personal/app/helpers"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authServ services.AuthService) AuthController {
	return &authController{
		authService: authServ,
	}
}

func (c *authController) Register(context *gin.Context) {
	var authCreateDTO dtos.AuthCreateDTO
	errDTO := context.ShouldBind(&authCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.authService.Register(authCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *authController) Login(context *gin.Context) {
	var authLoginDTO dtos.AuthLoginDTO
	errDTO := context.ShouldBind(&authLoginDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.authService.Login(authLoginDTO)
	if result == "" {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, res)
	}
}
