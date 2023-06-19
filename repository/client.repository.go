package repository

import (
	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
)

type ClientRepository interface {
	GetAllClient() ([]*models.Client, error)
	GetSingleClient(id string) (*models.Client, error)
	InsertClient(client models.Client) (string, error)
	UpdateClient(client models.Client, id string) error
	DeleteClient(id string) error
}

const (
	getAllClientQuery    = "SELECT id, name, token_sso, created_at, updated_at FROM client WHERE deleted_at is null"
	getSingleClientQuery = "SELECT id, name, token_sso, created_at, updated_at FROM client WHERE deleted_at is null AND id = $1"
)

func (r *Repository) GetAllClient() ([]*models.Client, error) {
	var clients []*models.Client

	err := r.db.Select(&clients, getAllClientQuery)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *Repository) GetSingleClient(id string) (*models.Client, error) {
	var clients []*models.Client

	err := r.db.Select(&clients, getSingleClientQuery, id)
	if err != nil {
		return nil, err
	}

	if len(clients) == 0 {
		return nil, nil
	}

	return clients[0], nil
}
