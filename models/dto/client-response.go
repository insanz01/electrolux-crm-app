package dto

type ClientResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	TokenSSO  string `json:"token_sso"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
