package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type ProductLineController interface {
	FindAll(c echo.Context) error
	// FindById(c echo.Context) error
	Insert(c echo.Context) error
}

type (
	productLineController struct {
		service services.ProductLineService
	}
)

func NewProductLineController(service services.ProductLineService) ProductLineController {
	return &productLineController{
		service: service,
	}
}

func (cc *productLineController) FindAll(c echo.Context) error {
	productLines, err := cc.service.FindAll(c)
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
		Data:    productLines,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (cc *productLineController) Insert(c echo.Context) error {
	productLineInsert := dto.ProductLineInsertRequest{}

	c.Bind(&productLineInsert)

	productLine, err := cc.service.Insert(c, productLineInsert)
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
		Data:    productLine,
	}

	return c.JSON(http.StatusOK, webResponse)
}
