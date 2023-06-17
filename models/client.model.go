package models

type Client struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	TokenSSO  string `db:"token_sso"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
