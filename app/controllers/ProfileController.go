package controllers

import (
	"log"
	"net/http"
	"personal/app/configs"
	"personal/app/helpers"
	"personal/app/models"
	"personal/app/services"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	Profile(context *gin.Context)
	Logout(context *gin.Context)
}

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileServ services.ProfileService) ProfileController {
	return &profileController{
		profileService: profileServ,
	}
}

func (c *profileController) Profile(context *gin.Context) {
	user_id, err := configs.ExtractTokenID(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := c.profileService.GetProfile(user_id)
	if (u == models.Auth{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", u)
		context.JSON(http.StatusOK, res)
	}
}

func (c *profileController) Logout(context *gin.Context) {
	au, err := configs.ExtractTokenID(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := c.profileService.Logout(au)
	if delErr != nil {
		log.Println(delErr)
		context.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	context.JSON(http.StatusOK, "Successfully logged out")
}
