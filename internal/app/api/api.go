package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kingofmidas/astrology-api/config"
	postgresPkg "github.com/kingofmidas/astrology-api/pkg/postgres"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	Router   *mux.Router
	PgClient *sqlx.DB
}

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

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(false)

	api := APIServer{
		Router:   router,
		PgClient: pgClient,
	}

	api.register(cfg)

	return api.start(cfg.ServerAddress)
}

func (s *APIServer) start(address string) error {
	srv := &http.Server{
		Addr:    address,
		Handler: s.Router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	logrus.Info("server starting ...")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen and serve: %v", err)
		}
	}()
	<-stop

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return srv.Shutdown(ctx)
}
