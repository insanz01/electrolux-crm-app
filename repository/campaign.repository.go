package repository

import (
	"errors"
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"github.com/jmoiron/sqlx"
)

type CampaignRepository interface {
	GetAllCampaign() ([]*models.Campaign, error)
	GetSingleCampaign(id string) (*models.Campaign, error)
	InsertCampaign(campaign models.Campaign) (string, error)
	CreateCampaignSummary(campaignSummary models.CampaignSummary) (string, error)
	CreateBatchCustomerCampaign(campaignSummaryId string, filters models.CampaignFilterProperties) (string, error)
}

const (
	getAllCampaignQuery         = "SELECT id, name, channel_account_id, client_id, city, count_repeat, is_repeated, is_scheduled, model_type, product_line, purchase_date, schedule_date, service_type, status, template_id FROM public.campaign"
	getSingleCampaignQuery      = "SELECT id, name, channel_account_id, client_id, city, count_repeat, is_repeated, is_scheduled, model_type, product_line, purchase_date, schedule_date, service_type, status, template_id FROM public.campaign WHERE public.campaign.id = $1"
	insertCampaignQuery         = "INSERT INTO public.campaign (name, channel_account_id, client_id, city, count_repeat, is_repeated, is_scheduled, model_type, product_line, purchase_date, schedule_date, service_type, status, template_id) VALUES (:name, :channel_account_id, :client_id, :city, :count_repeat, :is_repeated, :is_scheduled, :model_type, :product_line, :purcahse_data, :schedule_date, :service_type, :status, :template_id) returning id"
	insertCampaignSummaryQuery  = "INSERT INTO public.campaign_summary (campaign_id, failed_sent, success_sent, status) VALUES (:campaign_id, :failed_sent, :success_sent, :status) returning id"
	insertCampaignCustomerQuery = "INSERT INTO public.campaign_customer (summary_id, customer_id, sent_at, delivered_at, read_at) VALUES (:summary_id, :customer_id, :sent_at, :delivered_at, :read_at) returning id"

	getAllUserByCampaignQuery = "SELECT DISTINCT p.table_data_id FROM public.properties as p WHERE p.value IN (?)"
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

func (r *Repository) InsertCampaign(campaign models.Campaign) (string, error) {
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

func (r *Repository) CreateCampaignSummary(campaignSummary models.CampaignSummary) (string, error) {
	stmt, err := r.db.PrepareNamed(insertCampaignSummaryQuery)
	if err != nil {
		return "", errors.New("insert_campaign_summary" + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&campaignSummary).Scan(&uuid)
	if err != nil {
		return "", errors.New("insert_campaign_summary" + err.Error())
	}
	return uuid, nil
}

func (r *Repository) CreateBatchCustomerCampaign(campaignSummaryId string, filters models.CampaignFilterProperties) (string, error) {
	customerUUIDs, err := r.getAllCustomerByFilter(filters)
	if err != nil {
		return "", err
	}

	for _, customerId := range customerUUIDs {
		customerCampaign := models.CampaignCustomer{
			SummaryId:   campaignSummaryId,
			CustomerId:  customerId,
			SentAt:      nil,
			DeliveredAt: nil,
			ReadAt:      nil,
		}

		_, err := r.createCustomerCampaign(customerCampaign)
		if err != nil {
			fmt.Println(err)
		}
	}

	return "", nil
}

func (r *Repository) getAllCustomerByFilter(filters models.CampaignFilterProperties) ([]string, error) {
	var customers []*models.CustomerProperties

	// Persiapan query
	query, args, err := sqlx.In(getAllUserByCampaignQuery, filters.Filters)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	// Eksekusi query
	err = r.db.Select(&customers, query, args...)
	if err != nil {
		return nil, err
	}

	fmt.Println("query", query)

	uuidString := []string{}

	for _, customer := range customers {
		uuidString = append(uuidString, customer.TableDataID)
	}

	return uuidString, err
}

func (r *Repository) createCustomerCampaign(campaignCustomer models.CampaignCustomer) (string, error) {
	stmt, err := r.db.PrepareNamed(insertCampaignCustomerQuery)
	if err != nil {
		return "", errors.New("insert_campaign_customer" + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&campaignCustomer).Scan(&uuid)
	if err != nil {
		return "", errors.New("insert_campaign_customer" + err.Error())
	}
	return uuid, nil
}
