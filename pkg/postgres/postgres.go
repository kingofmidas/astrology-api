package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewClient(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logrus.Errorf("close db connection: %v", err)
	}
}
