package dto

type ReportDownloadRequest struct {
	ClientId     string `json:"client_id"`
	CampaignName string `json:"campaign_name"`
	ChannelId    string `json:"channel_id"`
	Status       string `json:"status"`
}

type ReportFilter struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type ReportProperties struct {
	Filters []*ReportFilter `json:"filters"`
}
