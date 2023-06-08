package dto

import (
	"time"
)

type CampaignInsertRequest struct {
	Name             string     `json:"name"`
	ChannelAccountId string     `json:"channel_account_id"`
	ClientId         string     `json:"client_id"`
	City             []string   `json:"city"`
	CountRepeat      *int       `json:"count_repeat"`
	NumOfOccurence   *int       `json:"num_of_occurence"`
	IsRepeated       bool       `json:"is_repeated"`
	IsScheduled      bool       `json:"is_scheduled"`
	ModelType        []string   `json:"model_type"`
	ProductLine      []string   `json:"product_line"`
	PurchaseDate     *time.Time `json:"purchase_date"`
	ScheduleDate     *time.Time `json:"schedule_date"`
	ServiceType      []string   `json:"service_type"`
	Status           string     `json:"status"`
	TemplateId       string     `json:"template_id"`
}
