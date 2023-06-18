package models

import "time"

type Client struct {
	Id        string     `db:"id"`
	Name      string     `db:"name"`
	TokenSSO  string     `db:"token_sso"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
