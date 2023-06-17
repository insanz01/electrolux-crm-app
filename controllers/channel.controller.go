package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type ChannelController interface {
	FindAll(c echo.Context) error
	FindAllAccount(c echo.Context) error
}

type channelController struct {
	channelService services.ChannelService
}

func NewChannelController(channelService services.ChannelService) ChannelController {
	return &channelController{
		channelService: channelService,
	}
}

func (cc *channelController) FindAll(c echo.Context) error {
	channel, err := cc.channelService.FindAll(c)
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
		Data:    channel,
	})
}

func (cc *channelController) FindAllAccount(c echo.Context) error {
	channelAccounts, err := cc.channelService.FindAllAccount(c)
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
		Data:    channelAccounts,
	})
}
