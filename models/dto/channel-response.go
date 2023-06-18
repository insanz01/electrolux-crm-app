package dto

import "time"

type ChannelResponse struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	UpdatedById   string     `json:"updated_by_id"`
	UpdatedByName string     `json:"updated_by_name"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type ChannelAccountResponse struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Token     string     `json:"token"`
	ClientId  string     `json:"client_id"`
	ChannelId string     `json:"channel_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
