package dto

import (
	"time"
)

type Campaign struct {
	Id                string     `json:"id"`
	Name              string     `json:"name"`
	ChannelAccountId  string     `json:"channel_account_id"`
	ClientId          string     `json:"client_id"`
	City              []string   `json:"city"`
	CountRepeat       *int       `json:"count_repeat"`
	IsRepeated        bool       `json:"is_repeated"`
	IsScheduled       bool       `json:"is_scheduled"`
	NumOfOccurence    *int       `json:"num_of_occurence"`
	RepeatType        string     `json:"repeat_type"`
	ModelType         []string   `json:"model_type"`
	ProductLine       []string   `json:"product_line"`
	PurchaseStartDate string     `json:"purchase_start_date"`
	PurchaseEndDate   string     `json:"purchase_end_date"`
	ScheduleDate      string     `json:"schedule_date"`
	ServiceType       []string   `json:"service_type"`
	HeaderParameter   []string   `json:"header_parameter"`
	BodyParameter     []string   `json:"body_parameter"`
	MediaParameter    string     `json:"media_parameter"`
	ButtonParameter   []string   `json:"button_parameter"`
	Status            string     `json:"status"`
	TemplateId        string     `json:"template_id"`
	TemplateName      string     `json:"template_name"`
	RejectionNote     *string    `json:"rejection_note"`
	SubmitByUserId    *string    `json:"submit_by_user_id"`
	SubmitByUserName  *string    `json:"submit_by_user_name"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
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

type CampaignCustomerDetail struct {
	PhoneNumber string `json:"phone_number"`
	State       string `json:"state"`
}

type CampaignCustomer struct {
	Id             string                 `json:"id"`
	SummaryId      string                 `json:"summary_id"`
	CustomerId     string                 `json:"customer_id"`
	CustomerDetail CampaignCustomerDetail `json:"customer_detail"`
	SentAt         *time.Time             `json:"sent_at"`
	DeliveredAt    *time.Time             `json:"delivered_at"`
	ReadAt         *time.Time             `json:"read_at"`
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
	Note       *string `json:"note,omitempty"`
}

type CampaignListResponse struct {
	ListData []string `json:"list"`
}
