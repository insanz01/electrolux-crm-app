package services

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type ClientService interface {
	FindAll(c echo.Context) ([]*dto.ClientResponse, error)
}

type clientService struct {
	repository *repository.Repository
}

func NewClientService(repository *repository.Repository) ClientService {
	return &clientService{
		repository: repository,
	}
}

func (cs *clientService) FindAll(c echo.Context) ([]*dto.ClientResponse, error) {
	clients, err := cs.repository.GetAllClient()
	if err != nil {
		return nil, err
	}

	clientResponses := []*dto.ClientResponse{}
	for _, client := range clients {
		clientResponses = append(clientResponses, &dto.ClientResponse{
			Id:        client.Id,
			Name:      client.Name,
			TokenSSO:  client.TokenSSO,
			CreatedAt: client.CreatedAt,
			UpdatedAt: client.UpdatedAt,
		})
	}

	return clientResponses, nil
}
