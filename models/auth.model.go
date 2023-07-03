package models

type AuthSSO struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Roles        []string `json:"roles"`
	Title        string   `json:"title"`
	Avatar       string   `json:"avatar"`
	Mobile       string   `json:"mobile"`
	Modules      []string `json:"modules"`
	ClientID     string   `json:"client_id"`
	Divisions    []string `json:"divisions"`
	LastName     string   `json:"last_name"`
	FirstName    string   `json:"first_name"`
	Applications []string `json:"applications"`
	DisplayName  string   `json:"display_name"`
}
