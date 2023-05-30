package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"usename" db:"username"`
	Password string `json:"password" db:"password"`
}
