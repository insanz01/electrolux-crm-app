package services

import (
	"errors"
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type CampaignService interface {
	FindAll(c echo.Context) (*dto.CampaignsResponse, error)
	FindById(c echo.Context, id string) (*dto.CampaignResponse, error)
	FindSummary(c echo.Context, id string) (*dto.SummaryCampaignResponse, error)
	FindCustomerBySummary(c echo.Context, summaryId string) (*dto.CampaignCustomerResponses, error)
	Insert(c echo.Context, campaign dto.CampaignParsedRequest) (*dto.CampaignResponse, error)
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
			Id:                campaign.Id,
			Name:              campaign.Name,
			City:              campaign.City,
			ChannelAccountId:  campaign.ChannelAccountId,
			ClientId:          campaign.ClientId,
			CountRepeat:       campaign.CountRepeat,
			NumOfOccurence:    &campaign.NumOfOccurence,
			RepeatType:        campaign.RepeatType,
			IsRepeated:        campaign.IsRepeated,
			IsScheduled:       campaign.IsScheduled,
			ModelType:         campaign.ModelType,
			ProductLine:       campaign.ProductLine,
			PurchaseStartDate: campaign.PurchaseStartDate.Format("2006-01-02"),
			PurchaseEndDate:   campaign.PurchaseEndDate.Format("2006-01-02"),
			ScheduleDate:      campaign.ScheduleDate.Format("2006-01-02"),
			ServiceType:       campaign.ServiceType,
			Status:            campaign.Status,
			TemplateId:        campaign.TemplateId,
		})
	}

	campaignResponse := dto.CampaignsResponse{
		Campaigns: allCampaigns,
	}

	return &campaignResponse, nil
}

func (r *campaignService) FindById(c echo.Context, id string) (*dto.CampaignResponse, error) {
	campaign, err := r.repository.GetSingleCampaign(id)
	if err != nil {
		return nil, err
	}

	if campaign == nil {
		return nil, errors.New("no data by id")
	}

	singleCampaign := dto.Campaign{
		Id:                campaign.Id,
		Name:              campaign.Name,
		City:              campaign.City,
		ChannelAccountId:  campaign.ChannelAccountId,
		ClientId:          campaign.ClientId,
		CountRepeat:       campaign.CountRepeat,
		NumOfOccurence:    &campaign.NumOfOccurence,
		IsRepeated:        campaign.IsRepeated,
		IsScheduled:       campaign.IsScheduled,
		RepeatType:        campaign.RepeatType,
		ModelType:         campaign.ModelType,
		ProductLine:       campaign.ProductLine,
		PurchaseStartDate: campaign.PurchaseStartDate.Format("2006-01-02"),
		PurchaseEndDate:   campaign.PurchaseEndDate.Format("2006-01-02"),
		ScheduleDate:      campaign.ScheduleDate.Format("2006-01-02"),
		ServiceType:       campaign.ServiceType,
		Status:            campaign.Status,
		TemplateId:        campaign.TemplateId,
	}

	return &dto.CampaignResponse{
		Campaign: singleCampaign,
	}, nil
}

func (r *campaignService) Insert(c echo.Context, campaignRequest dto.CampaignParsedRequest) (*dto.CampaignResponse, error) {
	campaignInsert := models.Campaign{
		Name:              campaignRequest.Name,
		ChannelAccountId:  campaignRequest.ChannelAccountId,
		ClientId:          campaignRequest.ClientId,
		City:              campaignRequest.City,
		CountRepeat:       campaignRequest.CountRepeat,
		NumOfOccurence:    *campaignRequest.NumOfOccurence,
		IsRepeated:        campaignRequest.IsRepeated,
		IsScheduled:       campaignRequest.IsScheduled,
		RepeatType:        campaignRequest.RepeatType,
		ModelType:         campaignRequest.ModelType,
		ProductLine:       campaignRequest.ProductLine,
		PurchaseStartDate: campaignRequest.PurchaseStartDate,
		PurchaseEndDate:   campaignRequest.PurchaseEndDate,
		ScheduleDate:      campaignRequest.ScheduleDate,
		ServiceType:       campaignRequest.ServiceType,
		Status:            campaignRequest.Status,
		TemplateId:        campaignRequest.TemplateId,
	}

	id, err := r.repository.InsertCampaign(campaignInsert)
	if err != nil {
		fmt.Println("error satu", err.Error())
		return nil, err
	}

	campaignSummary := models.CampaignSummary{
		CampaignId:  id,
		FailedSent:  "",
		SuccessSent: "",
		Status:      "INITIAL",
	}

	summaryId, err := r.repository.CreateCampaignSummary(campaignSummary)
	if err != nil {
		return nil, err
	}

	fmt.Println("summary id", summaryId)

	campaignFilter := models.CampaignFilterProperties{}

	// add filter list (product line, city/location, service type, model type, purchase date)
	if len(campaignInsert.ProductLine) > 0 {
		campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.ProductLine...)
	}

	if len(campaignInsert.City) > 0 {
		campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.City...)
	}

	if len(campaignInsert.ServiceType) > 0 {
		campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.ServiceType...)
	}

	if len(campaignInsert.ModelType) > 0 {
		campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.ModelType...)
	}

	if campaignInsert.PurchaseStartDate != nil {
		campaignFilter.DateRange.StartDate = campaignInsert.PurchaseStartDate
	}

	if campaignInsert.PurchaseEndDate != nil {
		// campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.PurchaseEndDate.Format("2006-01-02"))
		campaignFilter.DateRange.EndDate = campaignInsert.PurchaseEndDate
	}

	campaignCustomerId, err := r.repository.CreateBatchCustomerCampaign(summaryId, campaignFilter)
	if err != nil {
		fmt.Println("error dua", err.Error())
		return nil, err
	}

	fmt.Println("customer id", campaignCustomerId)

	campaignResponse := dto.Campaign{
		Id:                id,
		Name:              campaignRequest.Name,
		ChannelAccountId:  campaignRequest.ChannelAccountId,
		ClientId:          campaignRequest.ClientId,
		City:              campaignRequest.City,
		CountRepeat:       campaignRequest.CountRepeat,
		NumOfOccurence:    campaignRequest.NumOfOccurence,
		IsRepeated:        campaignRequest.IsRepeated,
		IsScheduled:       campaignRequest.IsScheduled,
		ModelType:         campaignRequest.ModelType,
		RepeatType:        campaignRequest.RepeatType,
		ProductLine:       campaignRequest.ProductLine,
		PurchaseStartDate: campaignRequest.PurchaseStartDate.Format("2006-01-02"),
		PurchaseEndDate:   campaignRequest.PurchaseEndDate.Format("2006-01-02"),
		ScheduleDate:      campaignRequest.ScheduleDate.Format("2006-01-02"),
		ServiceType:       campaignRequest.ServiceType,
		Status:            campaignRequest.Status,
		TemplateId:        campaignRequest.TemplateId,
	}

	return &dto.CampaignResponse{
		Campaign: campaignResponse,
	}, nil
}

func (cs *campaignService) FindSummary(c echo.Context, id string) (*dto.SummaryCampaignResponse, error) {
	summaryCampaign, err := cs.repository.GetSummaryByCampaignId(id)
	if err != nil {
		return nil, err
	}

	if summaryCampaign == nil {
		return nil, errors.New("no summary data")
	}

	summaryCampaignResp := dto.SummaryCampaign{
		Id:          summaryCampaign.Id,
		CampaignId:  summaryCampaign.CampaignId,
		FailedSent:  summaryCampaign.FailedSent,
		SuccessSent: summaryCampaign.SuccessSent,
		Status:      summaryCampaign.Status,
		CreatedAt:   summaryCampaign.CreatedAt,
		UpdatedAt:   summaryCampaign.UpdatedAt,
	}

	return &dto.SummaryCampaignResponse{
		SummaryCampaign: summaryCampaignResp,
	}, nil
}

func (cs *campaignService) FindCustomerBySummary(c echo.Context, summaryId string) (*dto.CampaignCustomerResponses, error) {
	customerCampaigns, err := cs.repository.GetCustomersBySummaryId(summaryId)
	if err != nil {
		return nil, err
	}

	if customerCampaigns == nil {
		return nil, errors.New("no customer data")
	}

	customerCampaignResp := []dto.CampaignCustomer{}
	for _, customer := range customerCampaigns {
		customerCampaignResp = append(customerCampaignResp, dto.CampaignCustomer{
			Id:          customer.Id,
			SummaryId:   customer.SummaryId,
			CustomerId:  customer.CustomerId,
			SentAt:      customer.SentAt,
			DeliveredAt: customer.DeliveredAt,
			ReadAt:      customer.ReadAt,
		})
	}

	return &dto.CampaignCustomerResponses{
		CampaignCustomers: customerCampaignResp,
	}, nil
}
