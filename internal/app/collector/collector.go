package collector

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kingofmidas/astrology-api/config"
	"github.com/kingofmidas/astrology-api/internal/pkg/repository"
	"github.com/kingofmidas/astrology-api/internal/pkg/services"
	postgresPkg "github.com/kingofmidas/astrology-api/pkg/postgres"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	pgClient, err := postgresPkg.NewClient(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("сonnect to database: %w", err)
	}
	defer pgClient.Close()

	imageRepository := repository.NewImageRepository(pgClient)
	collectorService := services.NewCollectorService(cfg.NasaAPIKey, imageRepository)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		collectorService.Collect(ctx)
		wg.Done()
	}()

	<-stop
	cancel()
	wg.Wait()

	return nil
}
