package app

import (
	"time"

	"github.com/ninja-dark/GoCloudCamp/internal/db"
	"github.com/ninja-dark/GoCloudCamp/internal/http"
)

type AppConfig struct{
	ShutdownTimeout time.Duration
}

type Config struct{
	App AppConfig
	Repository db.Config
	HTTP http.Config 
}