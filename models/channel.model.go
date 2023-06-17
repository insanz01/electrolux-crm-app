package models

type Channel struct {
	Id            string `db:"id"`
	Name          string `db:"name"`
	UpdatedById   string `db:"updated_by_id"`
	UpdatedByName string `db:"updated_by_name"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}

type ChannelAccount struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Token     string `db:"token"`
	ClientId  string `db:"client_id"`
	ChannelId string `db:"channel_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
