package services

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type CampaignService interface {
	FindAll(c echo.Context) (*dto.CampaignsResponse, error)
	FindById(c echo.Context, id string) error
}

type campaignService struct {
	repository *repository.Repository
}

func NewCampaignService(repository *repository.Repository) CampaignService {
	return &campaignService{
		repository: repository,
	}
}

func (r *campaignService) FindAll(c echo.Context) (*dto.CampaignsResponse, error) {
	campaigns, err := r.repository.GetAllCampaign()
	if err != nil {
		return nil, err
	}

	var allCampaigns []dto.Campaign

	for _, campaign := range campaigns {
		allCampaigns = append(allCampaigns, dto.Campaign{
			Id:               campaign.Id,
			Name:             campaign.Name,
			ChannelAccountId: campaign.ChannelAccountId,
			ClientId:         campaign.ClientId,
			CountRepeat:      campaign.CountRepeat,
			IsRepeated:       campaign.IsRepeated,
			IsScheduled:      campaign.IsScheduled,
			ModelType:        campaign.ModelType,
			ProductLine:      campaign.ProductLine,
			PurchaseDate:     campaign.PurchaseDate,
			ScheduleDate:     campaign.ScheduleDate,
			ServiceType:      campaign.ServiceType,
			Status:           campaign.Status,
			TemplateId:       campaign.TemplateId,
		})
	}

	campaignResponse := dto.CampaignsResponse{
		Campaigns: allCampaigns,
	}

	return &campaignResponse, nil
}

func (r *campaignService) FindById(c echo.Context, id string) error {
	return nil
}
