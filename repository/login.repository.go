package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/helpers"
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
)

type LoginRepository interface {
	CheckLogin(username, password string) (bool, error)
}

func (r *Repository) CheckLogin(username, password string) (bool, error) {
	obj := []models.User{}

	sqlStatement := "SELECT id, username, password FROM users WHERE username = $1"

	err := r.db.Select(&obj, sqlStatement, username)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	if len(obj) < 1 {
		return false, errors.New("tidak ada data")
	}

	match, err := helpers.CheckPasswordHash(password, obj[0].Password)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
