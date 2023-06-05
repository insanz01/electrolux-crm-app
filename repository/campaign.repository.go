package repository

import (
	"errors"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
)

type CampaignRepository interface {
	GetAllCampaign() ([]*models.Campaign, error)
	GetSingleCampaign(id string) (*models.Campaign, error)
	InsertCampaign(campaign dto.CampaignInsertRequest) (string, error)
}

const (
	getAllCampaignQuery    = "SELECT id, name, channel_account_id, client_id, city, count_repeat, is_repeated, is_scheduled, model_type, product_line, purchase_date, schedule_date, service_type, status, template_id FROM public.campaign"
	getSingleCampaignQuery = "SELECT id, name, channel_account_id, client_id, city, count_repeat, is_repeated, is_scheduled, model_type, product_line, purchase_date, schedule_date, service_type, status, template_id FROM public.campaign WHERE public.campaign.id = $1"
	insertCampaignQuery    = "INSERT INTO public.campaign (name, channel_account_id, client_id, city, count_repeat, is_repeated, is_scheduled, model_type, product_line, purchase_date, schedule_date, service_type, status, template_id) VALUES (:name, :channel_account_id, :client_id, :city, :count_repeat, :is_repeated, :is_scheduled, :model_type, :product_line, :purcahse_data, :schedule_date, :service_type, :status, :template_id) returning id"
)

func (r *Repository) GetAllCampaign() ([]*models.Campaign, error) {
	var campaigns []*models.Campaign

	err := r.db.Select(&campaigns, getAllCampaignQuery)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *Repository) GetSingleCampaign(id string) (*models.Campaign, error) {
	var campaign []*models.Campaign

	err := r.db.Select(&campaign, getSingleCampaignQuery, id)
	if err != nil {
		return nil, err
	}

	if len(campaign) == 0 {
		return nil, nil
	}

	return campaign[0], nil
}

func (r *Repository) InsertCampaign(campaign dto.CampaignInsertRequest) (string, error) {
	stmt, err := r.db.PrepareNamed(insertCampaignQuery)
	if err != nil {
		return "", errors.New("insert_excel_document" + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&campaign).Scan(&uuid)
	if err != nil {
		return "", errors.New("insert_excel_document" + err.Error())
	}
	return uuid, nil
}
