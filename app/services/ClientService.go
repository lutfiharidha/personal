package services

import (
	"log"
	"personal/app/dtos"
	"personal/app/models"
	"personal/app/repositories"

	"github.com/mashingan/smapping"
)

type ClientService interface {
	Insert(client dtos.ClientCreateDTO) models.Client
	Update(client dtos.ClientUpdateDTO) models.Client
	Restore(client dtos.ClientRestoreDTO) models.Client
	Delete(client models.Client) models.Client
	DeletePermanent(client models.Client)
	All() []models.Client
	Deleted() []models.Client
	FindByID(clientID string) models.Client
}

type clientService struct {
	clientRepository repositories.ClientRepository
}

func NewClientService(clientRepo repositories.ClientRepository) ClientService {
	return &clientService{
		clientRepository: clientRepo,
	}
}

func (service *clientService) Insert(client dtos.ClientCreateDTO) models.Client {
	clientuct := models.Client{}
	err := smapping.FillStruct(&clientuct, smapping.MapFields(&client))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.clientRepository.InsertClient(clientuct)
	return res
}

func (service *clientService) Update(client dtos.ClientUpdateDTO) models.Client {
	clientuct := models.Client{}
	err := smapping.FillStruct(&clientuct, smapping.MapFields(&client))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.clientRepository.UpdateClient(clientuct)
	return res
}

func (service *clientService) Delete(client models.Client) models.Client {
	return service.clientRepository.DeleteClient(client)
}

func (service *clientService) All() []models.Client {
	return service.clientRepository.AllClient()
}

func (service *clientService) FindByID(clientuctID string) models.Client {
	return service.clientRepository.FindClientByID(clientuctID)
}

func (service *clientService) Deleted() []models.Client {
	return service.clientRepository.DeletedClient()
}

func (service *clientService) Restore(client dtos.ClientRestoreDTO) models.Client {
	clientuct := models.Client{}
	err := smapping.FillStruct(&clientuct, smapping.MapFields(&client))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.clientRepository.RestoreClient(clientuct)
	return res
}

func (service *clientService) DeletePermanent(client models.Client) {
	service.clientRepository.DeletePermanentClient(client)
}
