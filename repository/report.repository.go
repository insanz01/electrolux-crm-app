package repository

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
)

type ReportRepository interface {
	GetAllReports() ([]*models.Report, error)
	GetAllReportsByFilter(filter dto.ReportFilter) ([]*models.Report, error)
	GetSingleReport(id string) (*models.Report, error)
}

const (
	getAllReportsQuery    = "SELECT campaign.id, campaign.name as campaign_name, channel_account.id as channel_account_id, channel_account.name as channel_name, client.id as client_id, client.name as division, campaign.created_at, campaign.status FROM campaign JOIN channel_account ON campaign.channel_account_id = channel_account.id JOIN client ON campaign.client_id = client.id"
	getSingleReportsQuery = "SELECT campaign.id, campaign.name as campaign_name, channel.name as channel_name, channel_account.id as channel_account_id, channel_account.name as channel_account_name, client.id as client_id, client.name as created_by, '' as division, '' as message_id, '' as contact, campaign.template_id, '' as template_name, '' as template_category, '' as template_language, '' as content_type, '' as approved_at, '' as approved_by, '' as wa_id, '' as reply_button, '' as reply_at, campaign.status as state, '' as invalid, campaign_customer.sent_at, campaign_customer.delivered_at, campaign_customer.read_at, campaign_summary.failed_sent as failed_at, '' as failed_detail, campaign.created_at, campaign_customer.customer_id FROM campaign JOIN channel_account ON campaign.channel_account_id = channel_account.id JOIN channel ON channel_account.channel_id = channel.id JOIN client ON campaign.client_id = client.id JOIN campaign_summary ON campaign.id = campaign_summary.campaign_id JOIN campaign_customer ON campaign_summary.id = campaign_customer.summary_id WHERE campaign.deleted_at is null AND campaign.id = $1"
)

func (r *Repository) GetAllReports() ([]*models.Report, error) {
	var reports []*models.Report

	err := r.db.Select(&reports, getAllReportsQuery)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *Repository) GetAllReportsByFilter(filter dto.ReportProperties) ([]*models.Report, error) {
	var reports []*models.Report

	finalQuery := getAllReportsQuery

	query := ""

	for _, f := range filter.Filters {
		switch f.Property {
		case "campaign_id":
			if f.Value != "" {
				query = fmt.Sprintf("%s AND campaign.id = '%s'", query, f.Value)
			}
		case "campaign_name":
			if f.Value != "" {
				query = fmt.Sprintf("%s AND campaign.name = '%s'", query, f.Value)
			}
		case "client_id":
			if f.Value != "" {
				query = fmt.Sprintf("%s AND client.id = '%s'", query, f.Value)
			}
		case "created_at":
			if f.Value != "" {
				query = fmt.Sprintf("%s AND DATE(campaign.created_at) = DATE('%s')", query, f.Value)
			}
		}
	}

	finalQuery = fmt.Sprintf("%s%s", finalQuery, query)

	err := r.db.Select(&reports, finalQuery)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *Repository) GetSingleReport(id string) ([]*models.DownloadReport, error) {
	var reports []*models.DownloadReport

	err := r.db.Select(&reports, getSingleReportsQuery, id)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
