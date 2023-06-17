package services

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type ChannelService interface {
	FindAll(c echo.Context) ([]*dto.ChannelResponse, error)
	FindAllAccount(c echo.Context) ([]*dto.ChannelAccountResponse, error)
}

type channelService struct {
	repository *repository.Repository
}

func NewChannelService(repository *repository.Repository) ChannelService {
	return &channelService{
		repository: repository,
	}
}

func (cs *channelService) FindAll(c echo.Context) ([]*dto.ChannelResponse, error) {
	channels, err := cs.repository.GetAllChannel()
	if err != nil {
		return nil, err
	}

	channelResponses := []*dto.ChannelResponse{}
	for _, channel := range channels {
		channelResponses = append(channelResponses, &dto.ChannelResponse{
			Id:            channel.Id,
			Name:          channel.Name,
			UpdatedById:   channel.UpdatedById,
			UpdatedByName: channel.UpdatedByName,
			CreatedAt:     channel.CreatedAt,
			UpdatedAt:     channel.UpdatedAt,
		})
	}

	return channelResponses, nil
}

func (cs *channelService) FindAllAccount(c echo.Context) ([]*dto.ChannelAccountResponse, error) {
	channelAccounts, err := cs.repository.GetAllChannelAccount()
	if err != nil {
		return nil, err
	}

	channelAccountResponses := []*dto.ChannelAccountResponse{}
	for _, channelAccount := range channelAccounts {
		channelAccountResponses = append(channelAccountResponses, &dto.ChannelAccountResponse{
			Id:        channelAccount.Id,
			Name:      channelAccount.Name,
			Token:     channelAccount.Token,
			ClientId:  channelAccount.ClientId,
			ChannelId: channelAccount.ChannelId,
			CreatedAt: channelAccount.CreatedAt,
			UpdatedAt: channelAccount.UpdatedAt,
		})
	}

	return channelAccountResponses, nil
}
