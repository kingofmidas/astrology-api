package main

import (
	"github.com/sirupsen/logrus"

	"github.com/kingofmidas/astrology-api/internal/app/api"
)

func main() {
	if err := api.Run(); err != nil {
		logrus.Fatalf("server shutdown: %v", err)
	}

	logrus.Infoln("server gracefully closed")
}
