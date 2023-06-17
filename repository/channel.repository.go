package repository

import "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"

type ChannelRepository interface {
	GetAllChannel() ([]*models.Channel, error)
	GetSingleChannel(id string) (*models.Channel, error)
	GetAllChannelAcount() ([]*models.ChannelAccount, error)
	GetSingleChannelAccount(id string) (*models.ChannelAccount, error)
	InsertChannel(channel models.Channel) (string, error)
	InsertChannelAccount(channelAccount models.ChannelAccount) (string, error)
	UpdateChannel(channel models.Channel, id string) error
	UpdateChannelAccount(ChannelAccount models.ChannelAccount, id string) error
	DeleteChannel(id string) error
	DeleteChannelAccount(id string) error
}

const (
	getAllChannelQuery        = "SELECT id, name, updated_by_id, updated_by_name, created_at, updated_at FROM channel"
	getAllChannelAccountQuery = "SELECT id, name, token, client_id, channel_id, created_at, updated_at FROM channel_account"
)

func (r *Repository) GetAllChannel() ([]*models.Channel, error) {
	var channels []*models.Channel

	err := r.db.Select(&channels, getAllChannelQuery)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (r *Repository) GetAllChannelAccount() ([]*models.ChannelAccount, error) {
	var channelAccounts []*models.ChannelAccount

	err := r.db.Select(&channelAccounts, getAllChannelAccountQuery)
	if err != nil {
		return nil, err
	}

	return channelAccounts, nil
}
