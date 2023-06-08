package controllers

import (
	"net/http"
	"strconv"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
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

// GetAllGiftClaims		godoc
// @Summary			Get All Gift Claims
// @Description		Get All Gift Claims from Db.
// @Produce			application/json
// @Tags			gift_claims
// @Success			200 {object} models.Response{}
// @Router			/gift_claims [get]
func (gc *giftController) FindAll(c echo.Context) error {

	giftClaimProperties := dto.GiftClaimProperties{}

	c.Bind(&giftClaimProperties)
	limitParam := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitParam)

	pageParam := c.QueryParam("page")
	page, _ := strconv.Atoi(pageParam)

	if limit == 0 {
		limit = 5
	}

	if page == 0 {
		page = 1
	}

	if giftClaimProperties.Properties == nil && giftClaimProperties.Filters == nil {
		giftClaims, err := gc.giftService.FindAll(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  0,
				Message: err.Error(),
				Data:    nil,
			})
		}

		webResponse := models.Response{
			Status:  1,
			Message: "success",
			Data:    giftClaims,
		}

		return c.JSON(http.StatusOK, webResponse)
	}

	giftClaims, err := gc.giftService.FindByPropsOrFilter(c, giftClaimProperties)
	if err != nil {

		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: err.Error(),
			Data:    nil,
		})
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    giftClaims,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// FindByIdGiftClaim 		godoc
// @Summary				Get Single Gift Claim by id.
// @Param				giftClaimId path string true "get gift claim by id"
// @Description			Return the gift claims whoes gift claim value mathes id.
// @Produce				application/json
// @Tags				gift_claims
// @Success				200 {object} models.Response{}
// @Router				/gift_claims/{giftClaimId} [get]
func (gc *giftController) FindById(c echo.Context) error {

	giftProperties := dto.GiftClaimProperties{}

	c.Bind(&giftProperties)

	giftClaimId := c.Param("id")

	if giftClaimId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  0,
			"message": "invalid parameter request",
			"data":    nil,
		})
	}

	if giftProperties.Properties == nil && giftProperties.Filters == nil {
		giftClaim, err := gc.giftService.FindById(c, giftClaimId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"status":  0,
				"message": "can't get data from db",
				"data":    nil,
			})
		}

		webResponse := models.Response{
			Status:  1,
			Message: "success",
			Data:    giftClaim,
		}

		return c.JSON(http.StatusOK, webResponse)
	}

	giftClaim, err := gc.giftService.FindByIdWithPropsOrFilter(c, giftProperties, giftClaimId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  0,
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    giftClaim,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// UpdateGiftClaim		godoc
// @Summary			Update gift_claims
// @Description		Update gift_claims data.
// @Param			giftClaimId path string true "update gift_claims by id"
// @Param			gift_claims body dto.GiftClaimUpdateRequest true  "Update gift_claims"
// @Tags			gift_claims
// @Produce			application/json
// @Success			200 {object} models.Response{}
// @Router			/gift_claims/{giftClaimId} [put]
func (gc *giftController) Update(c echo.Context) error {
	giftClaimRequest := dto.GiftClaimUpdateRequest{}

	giftClaimId := c.Param("id")
	if giftClaimId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  0,
			"message": "invalid request param",
			"data":    nil,
		})
	}

	err := c.Bind(&giftClaimRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  0,
			"message": "can't fetch data from request",
			"data":    nil,
		})
	}

	giftClaimResponse, err := gc.giftService.Update(c, giftClaimRequest, giftClaimId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  0,
			"message": "can't get data from db",
			"data":    nil,
		})
	}

	webResponse := models.Response{
		Status:  1,
		Message: "success",
		Data:    giftClaimResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// DeleteGiftClaims		godoc
// @Summary			Delete gift_claims
// @Description		Remove gift_claims data by id.
// @Produce			application/json
// @Tags			gift_claims
// @Success			200 {object} models.Response{}
// @Router			/gift_claims/{gift_claimId} [delete]
func (gc *giftController) Delete(c echo.Context) error {
	giftClaimId := c.Param("id")
	if giftClaimId == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  0,
			"message": "invalid request param",
			"data":    nil,
		})
	}

	err := gc.giftService.Delete(c, giftClaimId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  0,
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
