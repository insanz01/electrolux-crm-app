package repository

import "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"

type ReportRepository interface {
	GetAllReports() ([]*models.Report, error)
}

const (
	getAllReportsQuery = "SELECT campaign.id, campaign.name as campaign_name, channel_account.id as channel_account_id, channel_account.name as channel_name, client.id as client_id, client.name as division, campaign.created_at, campaign.status FROM campaign JOIN channel_account ON campaign.channel_account_id = channel_account.id JOIN client ON campaign.client_id = client.id"
)

func (r *Repository) GetAllReports() ([]*models.Report, error) {
	var reports []*models.Report

	err := r.db.Select(&reports, getAllReportsQuery)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
