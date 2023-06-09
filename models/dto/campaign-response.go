package dto

import (
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
