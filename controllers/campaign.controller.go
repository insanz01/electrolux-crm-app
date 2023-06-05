package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type CampaignController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Insert(c echo.Context) error
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

	campaign, err := cc.campaignService.Insert(c, campaignInsert)
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
