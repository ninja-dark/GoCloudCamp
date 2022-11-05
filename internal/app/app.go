package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ninja-dark/GoCloudCamp/internal/db/pgrepository"
	"github.com/ninja-dark/GoCloudCamp/internal/http"
	"github.com/ninja-dark/GoCloudCamp/internal/usecases"
)

type app struct {
	config *Config
	logic usecases.ServiceLogic

	http http.API
}

func (a *app) Serve() error{
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			a.Shutdown()
			break
		}
	}()
	return a.http.Serve()
}

func (a *app) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), a.config.App.ShutdownTimeout)
	defer cancel()
	a.http.Shutdown(ctx)
}

func New(ctx context.Context, config *Config)(App, error) {
	a := &app{
		config: config,
	}
	pgConfig, err := pgxpool.ParseConfig(config.Repository.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse conn string (%s): %w", config.Repository.ConnString, err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, pgConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	repository := pgrepository.New(pool)
	err = repository.InitSchema(ctx)
	if err != nil {
		return nil, err
	}
	a.http = http.New(&config.HTTP, a.logic)

	return a, nil
}