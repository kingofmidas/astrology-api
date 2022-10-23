package main

import (
	"github.com/kingofmidas/astrology-api/config"
	"github.com/kingofmidas/astrology-api/internal/app/migrate"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("load config: %v", err)
	}

	if err := migrate.Run(cfg.MigrationsPath, cfg.DatabaseURL); err != nil {
		logrus.Fatalf("run migrations: %v", err)
	}
}
