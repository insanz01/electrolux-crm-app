package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type CampaignService interface {
	FindAll(c echo.Context) (*dto.CampaignsResponse, error)
	FindById(c echo.Context, id string) (*dto.CampaignResponse, error)
	FindAllByFilter(c echo.Context, campaignProperties dto.CampaignProperties) (*dto.CampaignsResponse, error)
	FindSummary(c echo.Context, id string) (*dto.SummaryCampaignResponse, error)
	FindCustomerBySummary(c echo.Context, summaryId string) (*dto.CampaignCustomerResponses, error)
	Insert(c echo.Context, campaign dto.CampaignParsedRequest) (*dto.CampaignResponse, error)
	State(c echo.Context, statusRequest dto.StatusRequest) (*dto.StatusResponse, error)
	List(c echo.Context, property string) (*dto.CampaignListResponse, error)
	FilterCustomer(c echo.Context, sumaryId string, phoneCustomer dto.PhoneCustomerFilter) (*dto.CampaignCustomerResponses, error)
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
		// segment ini belum ditambahkan
		client, _ := r.repository.GetSingleClient(campaign.ClientId)

		clientName := ""
		if client != nil {
			clientName = client.Name
		}

		fmt.Println(clientName)
		// akhir dari segment client

		headerParameter, err := r.parseToArrayString(campaign.HeaderParameter)
		if err != nil {
			fmt.Println(err.Error())
		}

		bodyParameter, err := r.parseToArrayString(campaign.BodyParameter)
		if err != nil {
			fmt.Println(err.Error())
		}

		mediaParameter, err := r.parseToString(campaign.MediaParameter)
		if err != nil {
			fmt.Println(err.Error())
		}

		buttonParameter, err := r.parseToArrayString(campaign.ButtonParameter)
		if err != nil {
			fmt.Println(err.Error())
		}

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
			ScheduleDate:      campaign.ScheduleDate,
			ServiceType:       campaign.ServiceType,
			HeaderParameter:   headerParameter,
			BodyParameter:     bodyParameter,
			MediaParameter:    mediaParameter,
			ButtonParameter:   buttonParameter,
			Status:            campaign.Status,
			TemplateId:        campaign.TemplateId,
			TemplateName:      campaign.TemplateName,
			RejectionNote:     campaign.RejectionNote,
			SubmitByUserId:    campaign.SubmitByUserId,
			SubmitByUserName:  campaign.SubmitByUserName,
			CreatedAt:         campaign.CreatedAt,
			UpdatedAt:         campaign.UpdatedAt,
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
		return nil, errors.New("no data find by id")
	}

	// segment ini belum ditambahkan
	// client, _ := r.repository.GetSingleClient(campaign.ClientId)

	// clientName := ""
	// if client != nil {
	// 	clientName = client.Name
	// }

	// fmt.Println(clientName)
	// akhir dari segment client

	headerParameter, err := r.parseToArrayString(campaign.HeaderParameter)
	if err != nil {
		fmt.Println(err.Error())
	}

	bodyParameter, err := r.parseToArrayString(campaign.BodyParameter)
	if err != nil {
		fmt.Println(err.Error())
	}

	mediaParameter, err := r.parseToString(campaign.MediaParameter)
	if err != nil {
		fmt.Println(err.Error())
	}

	buttonParameter, err := r.parseToArrayString(campaign.ButtonParameter)
	if err != nil {
		fmt.Println(err.Error())
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
		ScheduleDate:      campaign.ScheduleDate,
		ServiceType:       campaign.ServiceType,
		HeaderParameter:   headerParameter,
		BodyParameter:     bodyParameter,
		MediaParameter:    mediaParameter,
		ButtonParameter:   buttonParameter,
		Status:            campaign.Status,
		TemplateId:        campaign.TemplateId,
		TemplateName:      campaign.TemplateName,
		RejectionNote:     campaign.RejectionNote,
		SubmitByUserId:    campaign.SubmitByUserId,
		SubmitByUserName:  campaign.SubmitByUserName,
		CreatedAt:         campaign.CreatedAt,
		UpdatedAt:         campaign.UpdatedAt,
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
		// HeaderParameter:   campaignRequest.HeaderParameter,
		// BodyParameter:     campaignRequest.BodyParameter,
		Status:           campaignRequest.Status,
		TemplateId:       campaignRequest.TemplateId,
		TemplateName:     campaignRequest.TemplateName,
		SubmitByUserId:   campaignRequest.SubmitByUserId,
		SubmitByUserName: campaignRequest.SubmitByUserName,
	}

	headerStruct, err := r.parseToJSONs(campaignRequest.HeaderParameter, "text")
	if err != nil {
		fmt.Println(err.Error())
	}

	bodyStruct, err := r.parseToJSONs(campaignRequest.BodyParameter, "text")
	if err != nil {
		fmt.Println(err.Error())
	}

	mediaStruct, err := r.parseToJSON(campaignRequest.MediaParameter, "file")
	if err != nil {
		fmt.Println(err.Error())
	}

	buttonStruct, err := r.parseToJSONs(campaignRequest.ButtonParameter, "file")
	if err != nil {
		fmt.Println(err.Error())
	}

	headerJson, err := json.Marshal(headerStruct)
	if err != nil {
		fmt.Println("data header" + err.Error())
	}

	bodyJson, err := json.Marshal(bodyStruct)
	if err != nil {
		fmt.Println("data body" + err.Error())
	}

	mediaJson, err := json.Marshal(mediaStruct)
	if err != nil {
		fmt.Println("data media" + err.Error())
	}

	buttonJson, err := json.Marshal(buttonStruct)
	if err != nil {
		fmt.Println("data button" + err.Error())
	}

	fmt.Println(headerJson)
	fmt.Println(bodyJson)
	fmt.Println(mediaJson)
	fmt.Println(buttonJson)

	campaignInsert.HeaderParameter = string(headerJson)
	campaignInsert.BodyParameter = string(bodyJson)
	campaignInsert.MediaParameter = string(mediaJson)
	campaignInsert.ButtonParameter = string(buttonJson)

	campaignInsert.Status = "WAITING APPROVAL"

	id, err := r.repository.InsertCampaign(campaignInsert)
	if err != nil {
		fmt.Println("error satu", err.Error())
		return nil, err
	}

	campaignSummary := models.CampaignSummary{
		CampaignId:  id,
		FailedSent:  "",
		SuccessSent: "",
		Status:      "WAITING APPROVAL",
	}

	summaryId, err := r.repository.CreateCampaignSummary(campaignSummary)
	if err != nil {
		return nil, err
	}

	fmt.Println("summary id", summaryId)

	campaignFilter := models.CampaignFilterProperties{}

	// add filter list (product line, city/location, service type, model type, purchase date)
	if len(campaignInsert.ProductLine) > 0 {
		for _, val := range campaignInsert.ProductLine {
			objFilter := models.ObjectFilter{
				Key:   "product_line",
				Value: val,
			}

			campaignFilter.Filters = append(campaignFilter.Filters, objFilter)
		}

		// campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.ProductLine...)
	}

	if len(campaignInsert.City) > 0 {
		// campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.City...)
		for _, val := range campaignInsert.City {
			objFilter := models.ObjectFilter{
				Key:   "city",
				Value: val,
			}

			campaignFilter.Filters = append(campaignFilter.Filters, objFilter)
		}
	}

	if len(campaignInsert.ServiceType) > 0 {
		for _, val := range campaignInsert.ServiceType {
			objFilter := models.ObjectFilter{
				Key:   "service_type",
				Value: val,
			}

			campaignFilter.Filters = append(campaignFilter.Filters, objFilter)
		}
		// campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.ServiceType...)
	}

	if len(campaignInsert.ModelType) > 0 {
		for _, val := range campaignInsert.ModelType {
			objFilter := models.ObjectFilter{
				Key:   "model_type",
				Value: val,
			}

			campaignFilter.Filters = append(campaignFilter.Filters, objFilter)
		}

		// campaignFilter.Filters = append(campaignFilter.Filters, campaignInsert.ModelType...)
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
		ScheduleDate:      campaignRequest.ScheduleDate,
		ServiceType:       campaignRequest.ServiceType,
		HeaderParameter:   campaignRequest.HeaderParameter,
		BodyParameter:     campaignRequest.BodyParameter,
		MediaParameter:    campaignRequest.MediaParameter,
		ButtonParameter:   campaignRequest.ButtonParameter,
		Status:            campaignRequest.Status,
		TemplateId:        campaignRequest.TemplateId,
		TemplateName:      campaignRequest.TemplateName,
		SubmitByUserId:    campaignRequest.SubmitByUserId,
		SubmitByUserName:  campaignRequest.SubmitByUserName,
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
		customerDetail, _ := cs.repository.GetSingle(customer.CustomerId)

		custMobileNo := ""
		state := "Invalid"

		for _, c := range customerDetail {
			if c.Key == "mobile_no" {
				custMobileNo = c.Value
			}
		}

		if custMobileNo != "" {

			custMobileNo = strings.Replace(custMobileNo, "+", "", -1)

			if strings.HasPrefix(custMobileNo, "628") {
				state = "Valid"
			}

			customerCampaignResp = append(customerCampaignResp, dto.CampaignCustomer{
				Id:         customer.Id,
				SummaryId:  customer.SummaryId,
				CustomerId: customer.CustomerId,
				CustomerDetail: dto.CampaignCustomerDetail{
					PhoneNumber: custMobileNo,
					State:       state,
				},
				SentAt:      customer.SentAt,
				DeliveredAt: customer.DeliveredAt,
				ReadAt:      customer.ReadAt,
			})
		}
	}

	return &dto.CampaignCustomerResponses{
		CampaignCustomers: customerCampaignResp,
	}, nil
}

func (cs *campaignService) validateState(state string) bool {
	if state == "ACCEPTED" || state == "REJECTED" || state == "FINISHED" || state == "ON GOING" || state == "WAITING APPROVAL" {
		return true
	}

	return false
}

func (cs *campaignService) State(c echo.Context, statusRequest dto.StatusRequest) (*dto.StatusResponse, error) {
	userInfo := c.Get("auth_token").(*models.AuthSSO)

	if userInfo.User.ID == nil {
		return nil, errors.New("invalid sso token data")
	}

	status := models.CampaignStatus{
		CampaignId: statusRequest.CampaignId,
		State:      statusRequest.State,
		Note:       statusRequest.Note,
		ApprovedBy: userInfo.User.Name,
	}

	if status.State == "STOP" {
		status.State = "FINISHED"
	}

	if status.State == "ACCEPTED" {
		status.State = "ON GOING"
	}

	if !cs.validateState(status.State) {
		return nil, errors.New("invalid state")
	}

	err := cs.repository.UpdateState(status)
	if err != nil {
		return nil, err
	}

	err = cs.repository.UpdateSummaryState(status)
	if err != nil {
		return nil, err
	}

	statusResponse := dto.StatusResponse{
		CampaignId: status.CampaignId,
		State:      status.State,
		Note:       statusRequest.Note,
	}

	return &statusResponse, nil
}

func (cs *campaignService) FindAllByFilter(c echo.Context, campaignProperties dto.CampaignProperties) (*dto.CampaignsResponse, error) {
	campaigns, err := cs.repository.GetAllCampaignWithFilter(campaignProperties)
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
			ScheduleDate:      campaign.ScheduleDate,
			ServiceType:       campaign.ServiceType,
			Status:            campaign.Status,
			TemplateId:        campaign.TemplateId,
			TemplateName:      campaign.TemplateName,
			SubmitByUserId:    campaign.SubmitByUserId,
			SubmitByUserName:  campaign.SubmitByUserName,
			RejectionNote:     campaign.RejectionNote,
			CreatedAt:         campaign.CreatedAt,
			UpdatedAt:         campaign.UpdatedAt,
		})
	}

	campaignResponse := dto.CampaignsResponse{
		Campaigns: allCampaigns,
	}

	return &campaignResponse, nil
}

func (cs *campaignService) List(c echo.Context, property string) (*dto.CampaignListResponse, error) {
	lists, err := cs.repository.GetAllCampaign()
	if err != nil {
		return nil, err
	}

	listResponse := []string{}
	unique := make(map[string]bool)
	for _, list := range lists {
		if !unique[list.Name] {
			unique[list.Name] = true
			listResponse = append(listResponse, list.Name)
		}
	}

	return &dto.CampaignListResponse{
		ListData: listResponse,
	}, nil
}

func (cs *campaignService) FilterCustomer(c echo.Context, summaryId string, phoneCustomer dto.PhoneCustomerFilter) (*dto.CampaignCustomerResponses, error) {
	customerCampaigns, err := cs.repository.GetCustomersBySummaryId(summaryId)
	if err != nil {
		return nil, err
	}

	if customerCampaigns == nil {
		return nil, errors.New("no customer data")
	}

	customerCampaignResp := []dto.CampaignCustomer{}
	for _, customer := range customerCampaigns {
		customerDetail, _ := cs.repository.GetSingle(customer.CustomerId)

		custMobileNo := ""
		state := "Invalid"

		for _, c := range customerDetail {
			if c.Key == "mobile_no" {
				custMobileNo = c.Value
			}
		}

		if custMobileNo != "" {

			custMobileNo = strings.Replace(custMobileNo, "+", "", -1)

			if strings.HasPrefix(custMobileNo, "628") {
				state = "Valid"
			}

			if custMobileNo == phoneCustomer.PhoneNumber {
				customerCampaignResp = append(customerCampaignResp, dto.CampaignCustomer{
					Id:         customer.Id,
					SummaryId:  customer.SummaryId,
					CustomerId: customer.CustomerId,
					CustomerDetail: dto.CampaignCustomerDetail{
						PhoneNumber: custMobileNo,
						State:       state,
					},
					SentAt:      customer.SentAt,
					DeliveredAt: customer.DeliveredAt,
					ReadAt:      customer.ReadAt,
				})
			}
		}
	}

	return &dto.CampaignCustomerResponses{
		CampaignCustomers: customerCampaignResp,
	}, nil
}

func (r *campaignService) categoryFile(ext string) string {
	switch strings.ToLower(ext) {
	default:
		return ""
	case "doc", "docx", "xls", "xlsx", "pdf", "csv", "ppt", "pptx":
		return "file/documents"
	case "jpg", "png", "jpeg", "gif", "bmp", "webp", "tiff", "svg", "ico":
		return "file/image"
	case "mp4", "avi", "mkv", "mov", "wmv", "flv", "webm", "m4v", "mpeg", "mpg":
		return "file/video"
	}
}

func (r *campaignService) isAcceptedFile(fileName string) (string, bool) {
	valid := false

	f := strings.Split(fileName, ".")

	fileExt := r.categoryFile(f[len(f)-1])

	if fileExt != "" {
		valid = true
	}

	return fileExt, valid
}

func (r *campaignService) parseToJSON(parameter, typeParameter string) (*models.CampaignJSONParameter, error) {
	if parameter == "" {
		return nil, errors.New("tidak ada data parameter")
	}

	typeData := typeParameter

	if typeParameter == "file" {
		tempType, isValid := r.isAcceptedFile(parameter)
		if isValid {
			typeData = tempType
		}
	}

	parameterStruct := models.CampaignJSONParameter{
		Type:   typeData,
		Number: 1,
		Value:  parameter,
	}

	return &parameterStruct, nil
}

func (r *campaignService) parseToJSONs(parameters []string, typeParameter string) ([]models.CampaignJSONParameter, error) {
	parameterStruct := []models.CampaignJSONParameter{}

	for idx, param := range parameters {
		typeData := typeParameter

		if typeParameter == "file" {
			tempType, isValid := r.isAcceptedFile(param)
			if isValid {
				typeData = tempType
			}
		}

		parameterStruct = append(parameterStruct, models.CampaignJSONParameter{
			Type:   typeData,
			Number: idx + 1,
			Value:  param,
		})
	}

	if len(parameterStruct) == 0 {
		return parameterStruct, errors.New("tidak ada data parameter")
	}

	return parameterStruct, nil
}

func (r *campaignService) parseToString(jsonString string) (string, error) {
	var paramVal models.CampaignJSONParameter

	err := json.Unmarshal([]byte(jsonString), &paramVal)
	if err != nil {
		return "", err
	}

	return paramVal.Value, nil
}

func (r *campaignService) parseToArrayString(jsonString string) ([]string, error) {
	arrString := []string{}
	var paramVal []models.CampaignJSONParameter

	err := json.Unmarshal([]byte(jsonString), &paramVal)
	if err != nil {
		return []string{}, err
	}

	for _, param := range paramVal {
		arrString = append(arrString, param.Value)
	}

	return arrString, nil
}
