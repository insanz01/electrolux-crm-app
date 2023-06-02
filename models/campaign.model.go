package models

import (
	"time"

	"github.com/lib/pq"
)

type Campaign struct {
	Id               string      `db:"id"`
	Name             string      `db:"name"`
	ChannelAccountId string      `db:"channel_account_id"`
	ClientId         string      `db:"client_id"`
	City             []string    `db:"city"`
	CountRepeat      *int        `db:"count_repeat"`
	IsRepeated       bool        `db:"is_repeated"`
	IsScheduled      bool        `db:"is_scheduled"`
	ModelType        []string    `db:"model_type"`
	ProductLine      []string    `db:"product_line"`
	PurchaseDate     pq.NullTime `db:"purchase_date"`
	ScheduleDate     pq.NullTime `db:"schedule_date"`
	ServiceType      []string    `db:"service_type"`
	Status           string      `db:"status"`
	TemplateId       string      `db:"template_id"`
}

type CampaignHistory struct {
	Id          string `db:"id"`
	CampaignId  string `db:"campaign_id"`
	FailedSent  string `db:"failed_sent"`
	SuccessSent string `db:"success_sent"`
	Status      string `db:"status"`
}

type CampaignCustomer struct {
	Id          string     `db:"id"`
	SummaryId   string     `db:"summary_id"`
	CustomerId  string     `db:"customer_id"`
	SentAt      *time.Time `db:"sent_at"`
	DeliveredAt *time.Time `db:"delivered_at"`
	ReadAt      *time.Time `db:"read_at"`
}
