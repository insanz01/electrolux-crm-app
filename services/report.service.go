package services

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
)

type ReportService interface {
	FindAll(c echo.Context) ([]*dto.ReportResponse, error)
}

type reportService struct {
	repository *repository.Repository
}

func NewReportService(repository *repository.Repository) ReportService {
	return &reportService{
		repository: repository,
	}
}

func (rs *reportService) FindAll(c echo.Context) ([]*dto.ReportResponse, error) {
	reports, err := rs.repository.GetAllReports()
	if err != nil {
		return nil, err
	}

	reportResponses := []*dto.ReportResponse{}
	for _, report := range reports {
		reportResponses = append(reportResponses, &dto.ReportResponse{
			Id:           report.Id,
			CampaignName: report.CampaignName,
			ChannelId:    report.ChannelId,
			ChannelName:  report.ChannelName,
			ClientId:     report.ClientId,
			Division:     report.Division,
			CreatedDate:  report.CreatedDate,
			Status:       report.Status,
		})
	}

	return reportResponses, nil
}
