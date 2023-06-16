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
