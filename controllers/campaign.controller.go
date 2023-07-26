package controllers

import (
	"fmt"
	"net/http"
	"time"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type CampaignController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Insert(c echo.Context) error
	Summary(c echo.Context) error
	Customer(c echo.Context) error
	Filter(c echo.Context) error
	Status(c echo.Context) error
	List(c echo.Context) error
	FilterCustomer(c echo.Context) error
}

type campaignController struct {
	campaignService services.CampaignService
}

func NewCampaignController(service services.CampaignService) CampaignController {
	return &campaignController{
		campaignService: service,
	}
}

func (cc *campaignController) FindAll(c echo.Context) error {
	campaigns, err := cc.campaignService.FindAll(c)
	if err != nil {
		webResponse := models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		}

		return c.JSON(http.StatusBadRequest, webResponse)
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    campaigns,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (cc *campaignController) FindById(c echo.Context) error {
	campaignId := c.Param("id")

	campaign, err := cc.campaignService.FindById(c, campaignId)
	if err != nil {
		webResponse := models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		}

		return c.JSON(http.StatusBadRequest, webResponse)
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    campaign,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (cc *campaignController) Insert(c echo.Context) error {
	campaignInsert := dto.CampaignInsertRequest{}

	if err := c.Bind(&campaignInsert); err != nil {
		webResponse := models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		}

		return c.JSON(http.StatusBadRequest, webResponse)
	}

	purchaseStartDate, err := time.Parse("2006-01-02", campaignInsert.PurchaseStartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid purchase start date value",
			Data:    nil,
		})
	}

	purchaseEndDate, err := time.Parse("2006-01-02", campaignInsert.PurchaseEndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid purchase end date value",
			Data:    nil,
		})
	}

	scheduledDate, err := time.Parse("2006-01-02", campaignInsert.ScheduleDate)
	if campaignInsert.ScheduleDate != "" {
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  0,
				Message: "invalid scheduled date value",
				Data:    nil,
			})
		}
	}

	fmt.Println(campaignInsert)

	userInfo := c.Get("auth_token").(*models.AuthSSO)

	fmt.Println(userInfo)
	if userInfo.User.ID == nil {
		return c.JSON(http.StatusUnauthorized, models.Response{
			Status:  0,
			Message: "invalid sso token data",
			Data:    nil,
		})
	}

	parsedRequest := dto.CampaignParsedRequest{
		Name:              campaignInsert.Name,
		ChannelAccountId:  campaignInsert.ChannelAccountId,
		ClientId:          campaignInsert.ClientId,
		City:              campaignInsert.City,
		CountRepeat:       campaignInsert.CountRepeat,
		NumOfOccurence:    &campaignInsert.NumOfOccurence,
		IsRepeated:        campaignInsert.IsRepeated,
		IsScheduled:       campaignInsert.IsScheduled,
		RepeatType:        campaignInsert.RepeatType,
		ModelType:         campaignInsert.ModelType,
		ProductLine:       campaignInsert.ProductLine,
		PurchaseStartDate: &purchaseStartDate,
		PurchaseEndDate:   &purchaseEndDate,
		ScheduleDate:      &scheduledDate,
		ServiceType:       campaignInsert.ServiceType,
		HeaderParameter:   campaignInsert.HeaderParameter,
		BodyParameter:     campaignInsert.BodyParameter,
		MediaParameter:    campaignInsert.MediaParameter,
		ButtonParameter:   campaignInsert.ButtonParameter,
		Status:            "WAITING APPROVAL",
		TemplateId:        campaignInsert.TemplateId,
		TemplateName:      campaignInsert.TemplateName,
		SubmitByUserId:    userInfo.User.ID,
		SubmitByUserName:  userInfo.User.Name,
	}

	campaign, err := cc.campaignService.Insert(c, parsedRequest)
	if err != nil {
		webResponse := models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		}

		return c.JSON(http.StatusInternalServerError, webResponse)
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    campaign,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (cc *campaignController) Summary(c echo.Context) error {
	campaignId := c.Param("campaign_id")

	if campaignId == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid parameter",
			Data:    nil,
		})
	}

	campaignSummary, err := cc.campaignService.FindSummary(c, campaignId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  1,
		Message: "success",
		Data:    campaignSummary,
	})
}

func (cc *campaignController) Customer(c echo.Context) error {
	summaryId := c.Param("summary_id")

	if summaryId == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid parameter",
			Data:    nil,
		})
	}

	customerCampaign, err := cc.campaignService.FindCustomerBySummary(c, summaryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  1,
		Message: "success",
		Data:    customerCampaign,
	})
}

func (cc *campaignController) Filter(c echo.Context) error {
	campaignFilterRequest := dto.CampaignProperties{}

	c.Bind(&campaignFilterRequest)

	if campaignFilterRequest.Filters == nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid body request",
			Data:    nil,
		})
	}

	campaigns, err := cc.campaignService.FindAllByFilter(c, campaignFilterRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  1,
		Message: "success",
		Data:    campaigns,
	})
}

func (cc *campaignController) Status(c echo.Context) error {
	campaignState := dto.StatusRequest{}

	c.Bind(&campaignState)

	if campaignState.CampaignId == "" || campaignState.State == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid body request",
			Data:    nil,
		})
	}

	campaigns, err := cc.campaignService.State(c, campaignState)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  1,
		Message: "success",
		Data:    campaigns,
	})
}

func (cc *campaignController) List(c echo.Context) error {
	listProperty := dto.ListProperty{}

	c.Bind(&listProperty)

	if listProperty.Property == nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid body request",
			Data:    nil,
		})
	}

	listData, err := cc.campaignService.List(c, *listProperty.Property)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  1,
		Message: "success",
		Data:    listData,
	})
}

func (cc *campaignController) FilterCustomer(c echo.Context) error {
	phoneCustomer := dto.PhoneCustomerFilter{}

	summaryId := c.Param("summary_id")

	c.Bind(&phoneCustomer)

	if phoneCustomer.PhoneNumber == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid body request",
			Data:    nil,
		})
	}

	if summaryId == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid parameter",
			Data:    nil,
		})
	}

	listData, err := cc.campaignService.FilterCustomer(c, summaryId, phoneCustomer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  1,
		Message: "success",
		Data:    listData,
	})
}
