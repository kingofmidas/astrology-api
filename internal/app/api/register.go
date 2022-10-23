package api

import (
	"github.com/kingofmidas/astrology-api/config"
	"github.com/kingofmidas/astrology-api/internal/app/api/handlers"
	"github.com/kingofmidas/astrology-api/internal/pkg/repository"
	"github.com/kingofmidas/astrology-api/internal/pkg/services"
)

func (s *APIServer) register(cfg *config.Config) {
	// Init repositories
	imageRepository := repository.NewImageRepository(s.PgClient)

	// Init services
	collectorService := services.NewCollectorService(cfg.NasaAPIKey, imageRepository)
	imageService := services.NewImageService(imageRepository, collectorService)

	// Init handlers
	imageHandler := handlers.NewImageHandler(imageService)

	imageHandler.InitRoutes(s.Router)
}
