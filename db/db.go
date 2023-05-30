package db

import (
	"database/sql"
	"fmt"
	"log"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sqlx.DB

// var err error

func Init() {
	conf := config.GetConfig()

	logger := logrus.WithField("func", "postgres.new").WithField("host", conf.DB_HOST).WithField("name", conf.DB_NAME).WithField("user", conf.DB_USERNAME).WithField("password", conf.DB_PASSWORD)

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME, conf.SSL_MODE)
	dbConn, err := sql.Open("postgres", url)
	if err != nil {
		logger.WithError(err).Error("error connecting to database")
		log.Fatal(err)
	}
	db = sqlx.NewDb(dbConn, "postgres")
	if err != nil {
		logger.WithError(err).Error("error connecting to database")
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.WithError(err).Error("error connecting to database")
		log.Fatal(err)
	}

	logger.Info("success")
}

func CreateCon() *sqlx.DB {
	return db
}
