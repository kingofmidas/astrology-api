package main

import (
	"github.com/kingofmidas/astrology-api/internal/app/collector"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := collector.Run(); err != nil {
		logrus.Fatalf("collector shutdown: %v", err)
	}

	logrus.Infoln("collector gracefully stoped")
}
