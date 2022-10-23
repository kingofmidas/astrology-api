package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment    string `env:"APP_ENV"`
	ServerAddress  string `env:"SERVER_ADDRESS"`
	MigrationsPath string `env:"MIGRATIONS_PATH"`
	DatabaseURL    string `env:"POSTGRES_URL"`
	NasaAPIKey     string `env:"NASA_API_KEY"`
}

var config Config

func Load() (*Config, error) {
	appEnv, ok := os.LookupEnv("APP_ENV")
	if !ok || appEnv == "local" {
		cleanenv.ReadConfig("local.env", &config)
	}
	err := cleanenv.ReadEnv(&config)
	return &config, err
}
