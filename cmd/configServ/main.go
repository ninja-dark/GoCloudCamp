package main

import (
	"context"
	"log"
	"time"

	"github.com/ninja-dark/GoCloudCamp/internal/app"
	"github.com/ninja-dark/GoCloudCamp/internal/db"
	"github.com/ninja-dark/GoCloudCamp/internal/http"
)

const (
	DefaultConnString      = "postgres://usr:pwd@localhost:5932/example?sslmode=disable"
	DefaultServeAddress    = "localhost:8080"
	DefaultShutdownTimeout = 20 * time.Second
	DefaultBasePath        = "/api/v1"
)

func main() {
	config := app.Config{
		App: app.AppConfig{
			ShutdownTimeout: DefaultShutdownTimeout,
		},
		Repository: db.Config{
			ConnString: DefaultConnString,
		},
		HTTP: http.Config{
			ServeAddress: DefaultServeAddress,
			BasePath:     DefaultBasePath,
		},
	}

	a, err := app.New(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Serve(); err != nil {
		log.Fatal(err)
	}
}
