package models

type Client struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	TokenSSO string `db:"token_sso"`
}
