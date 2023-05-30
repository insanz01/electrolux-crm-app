package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo"
)

type CustomerController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type (
	customerController struct {
		customerService services.CustomerService
	}
)

func NewCustomerController(service services.CustomerService) CustomerController {
	return &customerController{
		customerService: service,
	}
}

func (cc *customerController) FindAll(c echo.Context) error {
	customers, err := cc.customerService.FindAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error",
			"data":    nil,
		})
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    customers,
	})
}

func (cc *customerController) FindById(c echo.Context) error {

	customerId := c.Param("id")

	if customerId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid parameter request",
			"data":    nil,
		})
	}

	customer, err := cc.customerService.FindById(c, customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    customer,
	})
}

func (cc *customerController) Update(c echo.Context) error {
	customerRequest := dto.CustomerUpdateRequest{}

	customerId := c.Param("id")
	if customerId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request param",
			"data":    nil,
		})
	}

	err := c.Bind(&customerRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "can't fetch data from request",
			"data":    nil,
		})
	}

	customerResponse, err := cc.customerService.Update(c, customerRequest, customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    customerResponse,
	})
}

func (cc *customerController) Delete(c echo.Context) error {
	customerId := c.Param("id")
	if customerId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request param",
			"data":    nil,
		})
	}

	err := cc.customerService.Delete(c, customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "can't remove data from db",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    nil,
	})
}
