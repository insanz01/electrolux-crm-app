package services

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/repository"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type ReportService interface {
	FindAll(c echo.Context) ([]*dto.ReportResponse, error)
	FindAllByFilter(c echo.Context, filter dto.ReportProperties) ([]*dto.ReportResponse, error)
	Request(c echo.Context, request dto.ReportDownloadRequest) (*dto.ReportDownloadResponse, error)
	Download(c echo.Context, id string) (*dto.ReportDownloadResponse, error)
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

func (rs *reportService) FindAllByFilter(c echo.Context, filter dto.ReportProperties) ([]*dto.ReportResponse, error) {
	reports, err := rs.repository.GetAllReportsByFilter(filter)
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

func (rs *reportService) Request(c echo.Context, request dto.ReportDownloadRequest) (*dto.ReportDownloadResponse, error) {

	return nil, nil
}

func (rs *reportService) Download(c echo.Context, id string) (*dto.ReportDownloadResponse, error) {
	reports, err := rs.repository.GetSingleReport(id)
	if err != nil {
		return nil, err
	}

	theReports := []*models.DownloadReport{}
	for _, report := range reports {
		customerDetail, _ := rs.repository.GetSingle(report.CustomerId)

		custMobileNo := ""
		state := "true"

		for _, c := range customerDetail {
			if c.Key == "mobile_no" {
				custMobileNo = c.Value
			}
		}

		if custMobileNo != "" {
			report.Contact = custMobileNo
			state = "false"
		}

		report.Invalid = state

		theReports = append(theReports, report)
	}

	filename, err := rs.generateFile(id, theReports)
	if err != nil {
		return nil, err
	}

	req := c.Request()
	urlSchema := req.URL.Scheme
	if urlSchema == "" {
		urlSchema = "http"
	}

	url := fmt.Sprintf("%s://%s/assets/", urlSchema, req.Host)

	return &dto.ReportDownloadResponse{
		FilePath: url + filename,
	}, nil
}

func (rs *reportService) generateFile(id string, reports []*models.DownloadReport) (string, error) {
	fileName := fmt.Sprintf("uploads/%s.xlsx", id)

	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "ClientName")
	file.SetCellValue("Sheet1", "B1", "DivisionName")
	file.SetCellValue("Sheet1", "C1", "BroadcastId")
	file.SetCellValue("Sheet1", "D1", "ChannelAccount")
	file.SetCellValue("Sheet1", "E1", "CampaignName")
	file.SetCellValue("Sheet1", "F1", "MessageId")
	file.SetCellValue("Sheet1", "G1", "Channel")
	file.SetCellValue("Sheet1", "H1", "Contact")
	file.SetCellValue("Sheet1", "I1", "TemplateName")
	file.SetCellValue("Sheet1", "J1", "TemplateCategory")
	file.SetCellValue("Sheet1", "K1", "TemplateLanguage")
	file.SetCellValue("Sheet1", "L1", "ContentType")
	file.SetCellValue("Sheet1", "M1", "CreatedAt")
	file.SetCellValue("Sheet1", "N1", "CreatedBy")
	file.SetCellValue("Sheet1", "O1", "ApprovedAt")
	file.SetCellValue("Sheet1", "P1", "ApprovedBy")
	file.SetCellValue("Sheet1", "Q1", "WaID")
	file.SetCellValue("Sheet1", "R1", "ReplyButton")
	file.SetCellValue("Sheet1", "S1", "ReplyAt")
	file.SetCellValue("Sheet1", "T1", "State")
	file.SetCellValue("Sheet1", "U1", "Invalid")
	file.SetCellValue("Sheet1", "V1", "SentAt")
	file.SetCellValue("Sheet1", "W1", "DeliveredAt")
	file.SetCellValue("Sheet1", "X1", "ReadAt")
	file.SetCellValue("Sheet1", "Y1", "FailedAt")
	file.SetCellValue("Sheet1", "Z1", "FailedDetail")

	for idx, report := range reports {
		cellNumber := idx + 1

		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", cellNumber), report.CreatedBy)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", cellNumber), report.DivisionName)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", cellNumber), report.BroadcastId)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", cellNumber), report.ChannelAccountName)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", cellNumber), report.CampaignName)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", cellNumber), report.MessageId)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", cellNumber), report.ChannelName)
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", cellNumber), report.Contact)
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", cellNumber), report.TemplateName)
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", cellNumber), report.TemplateCategory)
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", cellNumber), report.TemplateLanguage)
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", cellNumber), report.ContentType)
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", cellNumber), report.CreatedDate)
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", cellNumber), report.CreatedBy)
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", cellNumber), report.ApprovedAt)
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", cellNumber), report.ApprovedBy)
		file.SetCellValue("Sheet1", fmt.Sprintf("Q%d", cellNumber), report.WAID)
		file.SetCellValue("Sheet1", fmt.Sprintf("R%d", cellNumber), report.ReplyButton)
		file.SetCellValue("Sheet1", fmt.Sprintf("S%d", cellNumber), report.ReplyAt)
		file.SetCellValue("Sheet1", fmt.Sprintf("T%d", cellNumber), report.State)
		file.SetCellValue("Sheet1", fmt.Sprintf("U%d", cellNumber), report.Invalid)
		file.SetCellValue("Sheet1", fmt.Sprintf("V%d", cellNumber), report.SentAt)
		file.SetCellValue("Sheet1", fmt.Sprintf("W%d", cellNumber), report.DeliveredAt)
		file.SetCellValue("Sheet1", fmt.Sprintf("X%d", cellNumber), report.ReadAt)
		file.SetCellValue("Sheet1", fmt.Sprintf("Y%d", cellNumber), report.FailedAt)
		file.SetCellValue("Sheet1", fmt.Sprintf("Z%d", cellNumber), report.FailedDetail)
	}

	if err := file.SaveAs(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
