package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type ReportController interface {
	FindAll(c echo.Context) error
	Filter(c echo.Context) error
	Request(c echo.Context) error
	Download(c echo.Context) error
}

type reportController struct {
	reportService services.ReportService
}

func NewReportController(service services.ReportService) ReportController {
	return &reportController{
		reportService: service,
	}
}

func (rc *reportController) FindAll(c echo.Context) error {
	reports, err := rc.reportService.FindAll(c)
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
		Data:    reports,
	})
}

func (rc *reportController) Filter(c echo.Context) error {
	reportFilterRequest := dto.ReportProperties{}

	c.Bind(&reportFilterRequest)

	if reportFilterRequest.Filters == nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid body request",
			Data:    nil,
		})
	}

	reports, err := rc.reportService.FindAllByFilter(c, reportFilterRequest)
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
		Data:    reports,
	})
}

func (rc *reportController) Request(c echo.Context) error {
	requestDownload := dto.ReportDownloadRequest{}

	c.Bind(&requestDownload)

	// report, err := rc.reportService

	return nil
}

func (rc *reportController) Download(c echo.Context) error {
	campaignId := c.Param("campaign_id")
	if campaignId == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  0,
			Message: "invalid parameter",
			Data:    nil,
		})
	}

	fileReport, err := rc.reportService.Download(c, campaignId)
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
		Data:    fileReport,
	})
}
