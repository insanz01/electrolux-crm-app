package controllers

import (
	"fmt"
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
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

// GetAllCustomers		godoc
// @Summary			Get All Customers
// @Description		Get All Customer from Db.
// @Produce			application/json
// @Tags			customer
// @Success			200 {object} models.Response{}
// @Router			/customers [get]
func (cc *customerController) FindAll(c echo.Context) error {
	customerProperties := dto.CustomerProperties{}

	c.Bind(&customerProperties)

	if customerProperties.Properties == nil {
		customers, err := cc.customerService.FindAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error",
				"data":    nil,
			})
			return nil
		}

		c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    customers,
		})
	} else {
		fmt.Println("Masuk ke sini")
		customers, err := cc.customerService.FindByPropsOrFilter(c, customerProperties)
		if err != nil {
			c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error",
				"data":    nil,
			})
			return nil
		}

		webResponse := models.Response{
			Status:  1,
			Message: "success",
			Data:    customers,
		}

		c.JSON(http.StatusOK, webResponse)
	}

	return nil
}

// FindByIdCustomer 		godoc
// @Summary				Get Single customer by id.
// @Param				customerId path string true "get customer by id"
// @Description			Return the customers whoes customer value mathes id.
// @Produce				application/json
// @Tags				customers
// @Success				200 {object} models.Response{}
// @Router				/customers/{customerId} [get]
func (cc *customerController) FindById(c echo.Context) error {

	customerProperties := dto.CustomerProperties{}

	c.Bind(&customerProperties)

	customerId := c.Param("id")

	if customerId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid parameter request",
			"data":    nil,
		})
	}

	if customerProperties.Properties == nil && customerProperties.Filters == nil {
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

	customer, err := cc.customerService.FindByIdWithPropsOrFilter(c, customerProperties, customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    customer,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// UpdateCustomer		godoc
// @Summary			Update customers
// @Description		Update customers data.
// @Param			id path string true "update customers by id"
// @Param			customers body dto.CustomerUpdateRequest true  "Update customers"
// @Tags			customers
// @Produce			application/json
// @Success			200 {object} models.Response{}
// @Router			/customers/{customerId} [put]
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

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    customerResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// DeleteCustomers		godoc
// @Summary			Delete customers
// @Description		Remove customers data by id.
// @Produce			application/json
// @Tags			customers
// @Success			200 {object} models.Response{}
// @Router			/customers/{customerId} [delete]
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

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, webResponse)
}
