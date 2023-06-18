package dto

import "time"

type ClientResponse struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	TokenSSO  string     `json:"token_sso"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
