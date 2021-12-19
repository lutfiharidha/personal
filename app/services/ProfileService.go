package services

import (
	"personal/app/models"
	"personal/app/repositories"
)

type ProfileService interface {
	GetProfile(uid string) models.Auth
	Logout(uid string) error
}

type profileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepo repositories.ProfileRepository) ProfileService {
	return &profileService{
		profileRepository: profileRepo,
	}
}

func (service *profileService) GetProfile(uid string) models.Auth {
	return service.profileRepository.Profile(uid)
}

func (service *profileService) Logout(uid string) error {
	return service.profileRepository.Logout(uid)
}
