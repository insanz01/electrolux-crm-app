package dto

type ReportDownloadRequest struct {
	ClientId     string `json:"client_id"`
	CampaignName string `json:"campaign_name"`
	ChannelId    string `json:"channel_id"`
	Status       string `json:"status"`
}
