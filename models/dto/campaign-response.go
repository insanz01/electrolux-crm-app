package dto

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	ChannelAccountId  uuid.UUID `json:"channel_account_id"`
	ClientId          uuid.UUID `json:"client_id"`
	City              []string  `json:"city"`
	CountRepeat       *int      `json:"count_repeat"`
	IsRepeated        bool      `json:"is_repeated"`
	IsScheduled       bool      `json:"is_scheduled"`
	NumOfOccurence    *int      `json:"num_of_occurence"`
	RepeatType        string    `json:"repeat_type"`
	ModelType         []string  `json:"model_type"`
	ProductLine       []string  `json:"product_line"`
	PurchaseStartDate string    `json:"purchase_start_date"`
	PurchaseEndDate   string    `json:"purchase_end_date"`
	ScheduleDate      string    `json:"schedule_date"`
	ServiceType       []string  `json:"service_type"`
	Status            string    `json:"status"`
	TemplateId        uuid.UUID `json:"template_id"`
}

type CampaignsResponse struct {
	Campaigns []Campaign `json:"campaigns"`
}

type CampaignResponse struct {
	Campaign Campaign `json:"campaign"`
}

type SummaryCampaign struct {
	Id          string     `json:"id"`
	CampaignId  string     `json:"campaign_id"`
	FailedSent  string     `json:"failed_sent"`
	SuccessSent string     `json:"success_sent"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type SummaryCampaignResponses struct {
	SummaryCampaigns []SummaryCampaign `json:"summary_campaigns"`
}

type SummaryCampaignResponse struct {
	SummaryCampaign SummaryCampaign `json:"summary_campaign"`
}

type CampaignCustomer struct {
	Id          string     `json:"id"`
	SummaryId   string     `json:"summary_id"`
	CustomerId  string     `json:"customer_id"`
	SentAt      *time.Time `json:"sent_at"`
	DeliveredAt *time.Time `json:"delivered_at"`
	ReadAt      *time.Time `json:"read_at"`
}

type CampaignCustomerResponses struct {
	CampaignCustomers []CampaignCustomer `json:"customer_campaigns"`
}

type CampaignCustomerResponse struct {
	CampaignCustomer CampaignCustomer `json:"customer_campaign"`
}

type StatusResponse struct {
	CampaignId string  `json:"campaign_id"`
	State      string  `json:"status"`
	Note       *string `json:"note"`
}
