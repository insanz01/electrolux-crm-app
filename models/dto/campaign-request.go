package dto

import (
	"time"

	"github.com/google/uuid"
)

type CampaignInsertRequest struct {
	Name              string    `json:"name"`
	ChannelAccountId  uuid.UUID `json:"channel_account_id"`
	ClientId          uuid.UUID `json:"client_id"`
	City              []string  `json:"city"`
	CountRepeat       *int      `json:"count_repeat"`
	NumOfOccurence    int       `json:"num_of_occurence"`
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

type CampaignParsedRequest struct {
	Name              string     `json:"name"`
	ChannelAccountId  uuid.UUID  `json:"channel_account_id"`
	ClientId          uuid.UUID  `json:"client_id"`
	City              []string   `json:"city"`
	CountRepeat       *int       `json:"count_repeat"`
	NumOfOccurence    *int       `json:"num_of_occurence"`
	IsRepeated        bool       `json:"is_repeated"`
	IsScheduled       bool       `json:"is_scheduled"`
	RepeatType        string     `json:"repeat_type"`
	ModelType         []string   `json:"model_type"`
	ProductLine       []string   `json:"product_line"`
	PurchaseStartDate *time.Time `json:"purchase_start_date"`
	PurchaseEndDate   *time.Time `json:"purchase_end_date"`
	ScheduleDate      *time.Time `json:"schedule_date"`
	ServiceType       []string   `json:"service_type"`
	Status            string     `json:"status"`
	TemplateId        uuid.UUID  `json:"template_id"`
}

type CampaignInsertV2Request struct {
	Name             string `json:"name"`
	ChannelAccountId string `json:"channel_account_id"`
	ClientId         string `json:"client_id"`
	City             string `json:"city"`
	CountRepeat      *int   `json:"count_repeat"`
	NumOfOccurence   *int   `json:"num_of_occurence"`
	IsRepeated       bool   `json:"is_repeated"`
	IsScheduled      bool   `json:"is_scheduled"`
	ModelType        string `json:"model_type"`
	ProductLine      string `json:"product_line"`
	PurchaseDate     string `json:"purchase_date"`
	ScheduleDate     string `json:"schedule_date"`
	ServiceType      string `json:"service_type"`
	Status           string `json:"status"`
	TemplateId       string `json:"template_id"`
}

type CampaignFilter struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type CampaignProperties struct {
	Target  string            `json:"target"`
	Filters []*CampaignFilter `json:"filters"`
}

type StatusRequest struct {
	State      string  `json:"state"`
	CampaignId string  `json:"campaign_id"`
	Note       *string `json:"note"`
}

type PhoneCustomerFilter struct {
	PhoneNumber string `query:"phone_number"`
}
