package repository

import "git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"

type ClientRepository interface {
	GetAllClient() ([]*models.Client, error)
	GetSingleClient(id string) (*models.Client, error)
	InsertClient(client models.Client) (string, error)
	UpdateClient(client models.Client, id string) error
	DeleteClient(id string) error
}

const (
	getAllClientQuery = "SELECT id, name, token_sso, created_at, updated_at FROM client"
)

func (r *Repository) GetAllClient() ([]*models.Client, error) {
	var clients []*models.Client

	err := r.db.Select(&clients, getAllClientQuery)
	if err != nil {
		return nil, err
	}

	return clients, nil
}
