package repository

import (
	"errors"
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models/dto"
	"github.com/jmoiron/sqlx"
)

type CampaignRepository interface {
	GetAllCampaign() ([]*models.Campaign, error)
	GetSingleCampaign(id string) (*models.Campaign, error)
	InsertCampaign(campaign models.Campaign) (string, error)
	CreateCampaignSummary(campaignSummary models.CampaignSummary) (string, error)
	CreateBatchCustomerCampaign(campaignSummaryId string, filters models.CampaignFilterProperties) (string, error)
	GetSummaryByCampaignId(id string) (*models.Campaign, error)
	GetCustomersBySummaryId(summaryId string) ([]*models.CampaignCustomer, error)
	GetAllCampaignWithFilter(filter dto.CampaignProperties) ([]*models.Campaign, error)
	UpdateState(status models.CampaignStatus) error
}

const (
	getAllCampaignQuery           = "SELECT id, name, channel_account_id, client_id, city, count_repeat, num_of_occurence, is_repeated, is_scheduled, model_type, product_line, purchase_start_date, purchase_end_date, repeat_type, schedule_date, service_type, status, template_id FROM public.campaign ORDER BY created_at DESC"
	getAllCampaignWithFilterQuery = "SELECT id, name, channel_account_id, client_id, city, count_repeat, num_of_occurence, is_repeated, is_scheduled, model_type, product_line, purchase_start_date, purchase_end_date, repeat_type, schedule_date, service_type, status, template_id FROM public.campaign WHERE deleted_at is NULL"
	getSingleCampaignQuery        = "SELECT id, name, channel_account_id, client_id, city, count_repeat, num_of_occurence, is_repeated, is_scheduled, model_type, product_line, purchase_start_date, purchase_end_date, repeat_type, schedule_date, service_type, status, template_id FROM public.campaign WHERE public.campaign.id = $1"
	insertCampaignQuery           = "INSERT INTO public.campaign (name, channel_account_id, client_id, city, count_repeat, num_of_occurence, is_repeated, is_scheduled, model_type, product_line, purchase_start_date, purchase_end_date, repeat_type, schedule_date, service_type, status, template_id) VALUES (:name, :channel_account_id, :client_id, :city, :count_repeat, :num_of_occurence, :is_repeated, :is_scheduled, :model_type, :product_line, :purchase_start_date, :purchase_end_date, :repeat_type, :schedule_date, :service_type, :status, :template_id) returning id"
	insertCampaignSummaryQuery    = "INSERT INTO public.campaign_summary (campaign_id, failed_sent, success_sent, status) VALUES (:campaign_id, :failed_sent, :success_sent, :status) returning id"
	insertCampaignCustomerQuery   = "INSERT INTO public.campaign_customer (summary_id, customer_id, sent_at, delivered_at, read_at) VALUES (:summary_id, :customer_id, :sent_at, :delivered_at, :read_at) returning id"

	getAllUserByCampaignQuery = "SELECT DISTINCT p.table_data_id FROM public.properties as p WHERE p.value IN (?)"

	getAllSummaryByCampaignIdQuery         = "SELECT s.id, s.campaign_id, s.failed_sent, s.success_sent, s.status, s.created_at, s.updated_at FROM campaign_summary s WHERE s.deleted_at is NULL AND s.campaign_id = $1"
	getAllCustomerCampaignBySummaryIdQuery = "SELECT cc.id, cc.summary_id, cc.customer_id, cc.sent_at, cc.delivered_at, cc.read_at FROM campaign_customer cc WHERE cc.summary_id = $1"

	updateCampaignStateQuery         = "UPDATE public.campaign SET status = :state, updated_at = NOW() WHERE id = :campaign_id"
	updateCampaignStateWithNoteQuery = "UPDATE public.campaign SET status = :state, rejection_note = :note, updated_at = NOW() WHERE id = :campaign_id"
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
		return "", errors.New("1_insert_campaign " + err.Error())
	}
	defer stmt.Close()

	var uuid string
	err = stmt.QueryRow(&campaign).Scan(&uuid)
	if err != nil {
		return "", errors.New("2_insert_campaign " + err.Error())
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

	finalQuery := getAllUserByCampaignQuery
	// filterDateQuery := "AND DATE(p.value) BETWEEN ? AND ?"

	// finalQuery = fmt.Sprintf("%s %s", finalQuery, filterDateQuery)
	// Persiapan query
	// query, args, err := sqlx.In(finalQuery, filters.Filters, filters.DateRange.StartDate, filters.DateRange.EndDate)
	query, args, err := sqlx.In(finalQuery, filters.Filters)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	fmt.Printf("filters: %v\n", filters)
	fmt.Println("query:", query)

	// Eksekusi query
	err = r.db.Select(&customers, query, args...)
	if err != nil {
		return nil, err
	}

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

func (r *Repository) GetSummaryByCampaignId(campaignId string) (*models.CampaignSummary, error) {
	var campaign []*models.CampaignSummary

	err := r.db.Select(&campaign, getAllSummaryByCampaignIdQuery, campaignId)
	if err != nil {
		return nil, err
	}

	if len(campaign) == 0 {
		return nil, nil
	}

	return campaign[0], nil
}

func (r *Repository) GetCustomersBySummaryId(summaryId string) ([]*models.CampaignCustomer, error) {
	var customerCampaigns []*models.CampaignCustomer

	err := r.db.Select(&customerCampaigns, getAllCustomerCampaignBySummaryIdQuery, summaryId)
	if err != nil {
		return nil, err
	}

	return customerCampaigns, nil
}

func (r *Repository) GetAllCampaignWithFilter(filter dto.CampaignProperties) ([]*models.Campaign, error) {
	var customerCampaigns []*models.Campaign

	finalQuery := getAllCampaignWithFilterQuery

	target := filter.Target

	query := ""

	for _, f := range filter.Filters {
		switch f.Property {
		case "name":
			query = fmt.Sprintf("%s AND name = '%s'", query, f.Value)
		case "channel_account_id":
			query = fmt.Sprintf("%s AND channel_account_id = '%s'", query, f.Value)
		case "status":
			query = fmt.Sprintf("%s AND status = '%s'", query, f.Value)
		case "submit_at":
			query = fmt.Sprintf("%s AND date(created_at) = date('%s')", query, f.Value)
		}
	}

	if target == "history" {
		query = fmt.Sprintf("%s AND (status = 'FINISHED' OR status = 'REJECTED')", query)
	} else {
		query = fmt.Sprintf("%s AND (status <> 'FINISHED' AND status <> 'REJECTED')", query)
	}

	finalQuery = fmt.Sprintf("%s%s", finalQuery, query)

	err := r.db.Select(&customerCampaigns, finalQuery)
	if err != nil {
		return nil, err
	}

	return customerCampaigns, nil
}

func (r *Repository) UpdateState(status models.CampaignStatus) error {
	finalQuery := updateCampaignStateQuery

	if status.State == "REJECTED" && status.Note != nil {
		finalQuery = updateCampaignStateWithNoteQuery
	}

	_, err := r.db.NamedExec(finalQuery, status)
	if err != nil {
		return err
	}

	return nil
}
