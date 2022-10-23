package migrate

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/sirupsen/logrus"
)

func Run(migrationsPath, databaseURL string) error {
	m, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		return fmt.Errorf("get migrate instance: %w", err)
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		logrus.Info("no change")
		return nil
	}

	if err != nil {
		return err
	}

	logrus.Info("success")
	return nil
}
