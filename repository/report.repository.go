package repository

import (
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
)

type ReportRepository interface {
	GetAllReports() ([]*models.Report, error)
	GetAllReportsByFilter(filter dto.ReportFilter) ([]*models.Report, error)
	GetSingleReport(id string) ([]*models.DownloadReport, error)
	GetRequestDownloadReport(request dto.ReportDownloadRequest) ([]*models.DownloadReport, error)
}

const (
	getAllReportsQuery    = "SELECT campaign.id, campaign.name as campaign_name, campaign.channel_account_id as channel_account_id, 'whatsapp' as channel_name, 'marketing' as division, campaign.client_id as client_id, campaign.client_id as client_name, campaign.created_at, campaign.status FROM campaign"
	getSingleReportsQuery = `SELECT 
		campaign.id, 
		campaign.name as campaign_name, 
		'nomor_wa' as channel_name, 
		campaign.channel_account_id as channel_account_id, 
		campaign.channel_account_id as channel_account_name, 
		campaign.client_id, 
		campaign.submit_by_user_name as created_by, 
		'marketing' as division, 
		campaign_customer.message_id as message_id, 
		'' as contact, 
		campaign.template_id, 
		campaign.template_name, 
		'' as template_category, 
		'id' as template_language, 
		'' as content_type, 
		campaign.approved_at, 
		campaign.approved_by as approved_by, 
		campaign_customer.wa_id as wa_id, 
		'' as reply_button, 
		'' as reply_at, 
		campaign.status as state, 
		'' as invalid, 
		campaign_customer.sent_at, 
		campaign_customer.delivered_at, 
		campaign_customer.read_at, 
		campaign_summary.failed_sent as failed_at, 
		'' as failed_detail, 
		campaign.created_at, 
		campaign_customer.customer_id 
		FROM campaign 
		JOIN campaign_summary 
		ON campaign.id = campaign_summary.campaign_id 
		JOIN campaign_customer 
		ON campaign_summary.id = campaign_customer.summary_id 
		WHERE campaign.deleted_at is null AND campaign.id = $1`
	getMultipleReportsQuery = `SELECT 
	campaign.id, 
	campaign.name as campaign_name, 
	'nomor_wa' as channel_name, 
	campaign.channel_account_id as channel_account_id, 
	campaign.channel_account_id as channel_account_name, 
	campaign.client_id, 
	campaign.submit_by_user_name as created_by, 
	'marketing' as division, 
	campaign_customer.message_id as message_id, 
	'' as contact, 
	campaign.template_id, 
	campaign.template_name, 
	'' as template_category, 
	'id' as template_language, 
	'' as content_type, 
	campaign.approved_at, 
	campaign.approved_by as approved_by, 
	campaign_customer.wa_id as wa_id, 
	'' as reply_button, 
	'' as reply_at, 
	campaign.status as state, 
	'' as invalid, 
	campaign_customer.sent_at, 
	campaign_customer.delivered_at, 
	campaign_customer.read_at, 
	campaign_summary.failed_sent as failed_at, 
	'' as failed_detail, 
	campaign.created_at, 
	campaign_customer.customer_id 
	FROM campaign 
	JOIN campaign_summary 
	ON campaign.id = campaign_summary.campaign_id 
	JOIN campaign_customer 
	ON campaign_summary.id = campaign_customer.summary_id 
	WHERE campaign.deleted_at is null`
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

func (r *Repository) GetRequestDownloadReport(request dto.ReportDownloadRequest) ([]*models.DownloadReport, error) {
	var reports []*models.DownloadReport

	finalQuery := getMultipleReportsQuery

	query := ""

	if request.CampaignName != "" {
		query = fmt.Sprintf("%s AND campaign.name = '%s'", query, request.CampaignName)
	}

	if request.ChannelId != "" {
		query = fmt.Sprintf("%s AND campaign.channel_account_id = '%s'", query, request.ChannelId)
	}

	if request.ClientId != "" {
		query = fmt.Sprintf("%s AND campaign.client_id = '%s'", query, request.ClientId)
	}

	if request.Status != "" {
		query = fmt.Sprintf("%s AND campaign.status = '%s'", query, request.Status)
	}

	finalQuery = fmt.Sprintf("%s%s", finalQuery, query)

	err := r.db.Select(&reports, finalQuery)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
