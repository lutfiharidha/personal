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
	authuct := models.Auth{}
	err := smapping.FillStruct(&authuct, smapping.MapFields(&auth))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.authRepository.RegisterAuth(authuct)
	return res
}

func (service *authService) Login(auth dtos.AuthLoginDTO) string {

	authuct := models.Auth{}
	err := smapping.FillStruct(&authuct, smapping.MapFields(&auth))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res, _ := service.authRepository.LoginCheck(authuct.Email, authuct.Password)
	if err != nil {
		return ""
	}
	return res

}
