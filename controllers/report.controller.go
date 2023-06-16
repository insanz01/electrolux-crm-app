package controllers

import (
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/services"
	"github.com/labstack/echo/v4"
)

type ReportController interface {
	FindAll(c echo.Context) error
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
