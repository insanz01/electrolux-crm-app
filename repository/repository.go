package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	Database interface {
		Select(dest interface{}, query string, args ...interface{}) error
		NamedExec(query string, arg interface{}) (sql.Result, error)
		PrepareNamed(query string) (*sqlx.NamedStmt, error)
	}
	Repository struct {
		db Database
	}
	Repo interface {
		LoginRepository
		CustomerRepository
		GiftRepository
		ChannelRepository
		CampaignRepository
	}
)

func New(db Database) *Repository {
	return &Repository{db: db}
}
