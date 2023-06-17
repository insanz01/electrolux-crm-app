package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type ClientController interface {
	FindAll(c echo.Context) error
}

type clientController struct {
	clientService services.ClientService
}

func NewClientController(clientService services.ClientService) ClientController {
	return &clientController{
		clientService: clientService,
	}
}

func (cc *clientController) FindAll(c echo.Context) error {
	clients, err := cc.clientService.FindAll(c)
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
		Data:    clients,
	})
}
