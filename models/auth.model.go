package models

type AuthSSO struct {
	AccessToken  string `json:"access_token"`
	ClientID     string `json:"client_id"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	User         struct {
		ID           *string  `json:"id"`
		Name         *string  `json:"name"`
		Email        string   `json:"email"`
		Phone        any      `json:"phone"`
		Roles        []string `json:"roles"`
		Title        any      `json:"title"`
		Avatar       any      `json:"avatar"`
		Mobile       any      `json:"mobile"`
		Modules      []string `json:"modules"`
		ClientID     string   `json:"client_id"`
		Divisions    []string `json:"divisions"`
		LastName     any      `json:"last_name"`
		FirstName    any      `json:"first_name"`
		Applications []string `json:"applications"`
		DisplayName  string   `json:"display_name"`
	} `json:"user"`
}
