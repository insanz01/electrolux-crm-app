package repository

import "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"

type ChannelRepository interface {
	GetAllChannel() ([]*models.Channel, error)
	GetSingleChannel(id string) (*models.Channel, error)
	GetAllChannelAcount() ([]models.ChannelAccount, error)
	GetSingleChannelAccount(id string) (*models.ChannelAccount, error)
	InsertChannel(channel models.Channel) (string, error)
	InsertChannelAccount(channelAccount models.ChannelAccount) (string, error)
	UpdateChannel(channel models.Channel, id string) error
	UpdateChannelAccount(ChannelAccount models.ChannelAccount, id string) error
	DeleteChannel(id string) error
	DeleteChannelAccount(id string) error
}
