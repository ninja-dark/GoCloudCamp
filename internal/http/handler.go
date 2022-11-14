package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("error here", err)
		return
	}
	var m enticonfig.Config
	json.Unmarshal(b, &m)
	if err := a.logic.CreateServ(m); err != nil{
		writeResponse(w, http.StatusBadRequest, fmt.Sprintf(`failed to parse user's id: %s`, err))
	}

	writeJsonResponse(w, http.StatusCreated, "Uploaded")
	
} 

func (a *api) GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	name := chi.URLParam(r, "service")
	s := a.logic
	conf, _ := s.GetConfig(ctx, name)

	writeJsonResponse(w, http.StatusOK, conf)
}

func (a *api) DeleteConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	name := chi.URLParam(r, "service")
	s := a.logic
	if err := s.DeleteConfig(ctx, name); err != nil{
		writeJsonResponse(w, http.StatusBadRequest, err)
	}

	writeJsonResponse(w, http.StatusOK, nil)

}

func (a *api) Serve() error {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		writeJsonResponse(w, http.StatusAccepted, "hello")
	})
	apiRouter.Post("/config", a.CreateService)
	apiRouter.Get("/config", a.GetConfig)
	apiRouter.Delete("/config", a.DeleteConfig)
	
	
	return http.ListenAndServe(a.config.ServeAddress, apiRouter)
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
