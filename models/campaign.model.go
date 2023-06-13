package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Campaign struct {
	Id                string         `db:"id"`
	Name              string         `db:"name"`
	ChannelAccountId  uuid.UUID      `db:"channel_account_id"`
	ClientId          uuid.UUID      `db:"client_id"`
	City              pq.StringArray `db:"city"`
	CountRepeat       *int           `db:"count_repeat"`
	NumOfOccurence    *int           `db:"num_of_occurence"`
	IsRepeated        bool           `db:"is_repeated"`
	IsScheduled       bool           `db:"is_scheduled"`
	RepeatType        string         `db:"repeat_type"`
	ModelType         pq.StringArray `db:"model_type"`
	ProductLine       pq.StringArray `db:"product_line"`
	PurchaseStartDate *time.Time     `db:"purchase_start_date"`
	PurchaseEndDate   *time.Time     `db:"purchase_end_date"`
	ScheduleDate      *time.Time     `db:"schedule_date"`
	ServiceType       pq.StringArray `db:"service_type"`
	Status            string         `db:"status"`
	TemplateId        uuid.UUID      `db:"template_id"`
}

type CampaignSummary struct {
	Id          string     `db:"id"`
	CampaignId  string     `db:"campaign_id"`
	FailedSent  string     `db:"failed_sent"`
	SuccessSent string     `db:"success_sent"`
	Status      string     `db:"status"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type CampaignCustomer struct {
	Id          string     `db:"id"`
	SummaryId   string     `db:"summary_id"`
	CustomerId  string     `db:"customer_id"`
	SentAt      *time.Time `db:"sent_at"`
	DeliveredAt *time.Time `db:"delivered_at"`
	ReadAt      *time.Time `db:"read_at"`
}

type CampaignDateRange struct {
	StartDate *time.Time `db:"purchase_start_date"`
	EndDate   *time.Time `db:"purchase_end_date"`
}

type CampaignFilterProperties struct {
	Filters   []string          `db:"filters" json:"filters"`
	DateRange CampaignDateRange `db:"daterange" json:"daterange"`
}
