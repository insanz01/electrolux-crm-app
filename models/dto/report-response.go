package dto

type ReportResponse struct {
	Id           string `json:"id"` // id campaign
	CampaignName string `json:"campaign_name"`
	ChannelId    string `json:"channel_account_id"`
	ChannelName  string `json:"channel_name"`
	ClientId     string `json:"client_id"`
	Division     string `json:"division"` // client
	CreatedDate  string `json:"created_at"`
	Status       string `json:"status"` // campaign status
}

type ReportResponses struct {
	ReportResponses []ReportResponse `json:"reports"`
}
