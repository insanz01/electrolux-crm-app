package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type GiftController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type (
	giftController struct {
		giftService services.GiftService
	}
)

func NewGiftController(service services.GiftService) GiftController {
	return &giftController{
		giftService: service,
	}
}

func (gc *giftController) FindAll(c echo.Context) error {

	giftClaimProperties := dto.GiftClaimProperties{}

	c.Bind(&giftClaimProperties)

	if giftClaimProperties.Properties == nil {
		giftClaims, err := gc.giftService.FindAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error",
				"data":    nil,
			})
			return nil
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    giftClaims,
		})
	}

	giftClaims, err := gc.giftService.FindAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error",
			"data":    nil,
		})
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    giftClaims,
	})
}

func (gc *giftController) FindById(c echo.Context) error {

	giftClaimId := c.Param("id")

	if giftClaimId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid parameter request",
			"data":    nil,
		})
	}

	giftClaim, err := gc.giftService.FindById(c, giftClaimId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    giftClaim,
	})
}

func (gc *giftController) Update(c echo.Context) error {
	giftClaimRequest := dto.GiftClaimUpdateRequest{}

	giftClaimId := c.Param("id")
	if giftClaimId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request param",
			"data":    nil,
		})
	}

	err := c.Bind(&giftClaimRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "can't fetch data from request",
			"data":    nil,
		})
	}

	giftClaimResponse, err := gc.giftService.Update(c, giftClaimRequest, giftClaimId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    giftClaimResponse,
	})
}

func (gc *giftController) Delete(c echo.Context) error {
	giftClaimId := c.Param("id")
	if giftClaimId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request param",
			"data":    nil,
		})
	}

	err := gc.giftService.Delete(c, giftClaimId)
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
