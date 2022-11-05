package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	enticonfig "github.com/ninja-dark/GoCloudCamp/internal/entiConfig"
	"github.com/ninja-dark/GoCloudCamp/internal/usecases"
)

type api struct {
	config *Config
	logic  usecases.ServiceLogic
	server *http.Server
}

func (a *api) CreateService(w http.ResponseWriter, r *http.Request) {
	s := &enticonfig.Config{}

	if err := json.NewDecoder(r.Body).Decode(s); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	writeJsonResponse(w, http.StatusCreated, s)
}

func (a *api) Serve() error {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		writeJsonResponse(w, http.StatusAccepted, "hello")
	})
	apiRouter.Post("/config", http.HandlerFunc(a.CreateService))
	a.server = &http.Server{Addr: a.config.ServeAddress, Handler: apiRouter}

	return a.server.ListenAndServe()
}

func (a *api) Shutdown(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
}

func New(
	config *Config,
	logic usecases.ServiceLogic) API {

	return &api{
		config: config,
		logic:  logic,
	}
}
