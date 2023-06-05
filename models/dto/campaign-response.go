package dto

import "github.com/lib/pq"

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

type CampaignsResponse struct {
	Campaigns []Campaign `json:"campaigns"`
}

type CampaignResponse struct {
	Campaign Campaign `json:"campaign"`
}