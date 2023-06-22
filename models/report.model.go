package models

type Report struct {
	Id           string `db:"id"` // id campaign
	CampaignName string `db:"campaign_name"`
	ChannelId    string `db:"channel_account_id"`
	ChannelName  string `db:"channel_name"`
	ClientId     string `db:"client_id"`
	Division     string `db:"division"` // client
	CreatedDate  string `db:"created_at"`
	Status       string `db:"status"` // campaign status
}

type DownloadReport struct {
	Id                 string  `db:"id"` // id campaign
	CampaignName       string  `db:"campaign_name"`
	ChannelId          string  `db:"channel_account_id"`
	ChannelName        string  `db:"channel_name"`
	ChannelAccountName string  `db:"channel_account_name"`
	CustomerId         string  `db:"customer_id"`
	Contact            string  `db:"contact"`
	DivisionName       string  `db:"division_name"`     // new
	MessageId          string  `db:"message_id"`        // new
	BroadcastId        string  `db:"broadcast_id"`      // new
	TemplateId         string  `db:"template_id"`       // new
	TemplateName       string  `db:"template_name"`     // new
	TemplateCategory   string  `db:"template_category"` // new
	TemplateLanguage   string  `db:"template_language"` // new
	ContentType        string  `db:"content_type"`      // new
	CreatedBy          string  `db:"created_by"`        // new
	ApprovedAt         string  `db:"approved_at"`       // new
	ApprovedBy         string  `db:"approved_by"`       // new
	WAID               string  `db:"wa_id"`             // new
	ReplyButton        string  `db:"reply_button"`      // new
	ReplyAt            string  `db:"reply_at"`          // new
	State              string  `db:"state"`             // new
	Invalid            string  `db:"invalid"`           // new
	SentAt             *string `db:"sent_at"`           // new
	DeliveredAt        *string `db:"delivered_at"`      // new
	ReadAt             *string `db:"read_at"`           // new
	FailedAt           string  `db:"failed_at"`         // new
	FailedDetail       string  `db:"failed_detail"`     // new
	ClientId           string  `db:"client_id"`
	Division           string  `db:"division"` // client
	CreatedDate        string  `db:"created_at"`
	Status             string  `db:"status"` // campaign status
}
