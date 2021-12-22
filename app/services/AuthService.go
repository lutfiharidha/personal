package services

import (
	"log"
	"personal/app/dtos"
	"personal/app/models"
	"personal/app/repositories"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Register(auth dtos.AuthCreateDTO) models.Auth
	Login(auth dtos.AuthLoginDTO) string
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepo,
	}
}

func (service *authService) Register(auth dtos.AuthCreateDTO) models.Auth {
	newAuth := models.Auth{}
	err := smapping.FillStruct(&newAuth, smapping.MapFields(&auth))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.authRepository.RegisterAuth(newAuth)
	return res
}

func (service *authService) Login(auth dtos.AuthLoginDTO) string {

	newAuth := models.Auth{}
	err := smapping.FillStruct(&newAuth, smapping.MapFields(&auth))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res, _ := service.authRepository.LoginCheck(newAuth.Email, newAuth.Password)
	if err != nil {
		return ""
	}
	return res

}
