package http

import (
	"context"
	"encoding/json"
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

	ctx := context.Background()
	b, _ := ioutil.ReadAll(r.Body)
	var rawMessage []json.RawMessage
	err := json.Unmarshal(b, &rawMessage)
	if err != nil {
		log.Fatal(err)
		return 
	}
	var name string
	type conf2 enticonfig.Config
	var c conf2
	err = json.Unmarshal(rawMessage[0], &c)
	if err != nil {
		log.Fatal(err)
		return 
	}
	name = c.Service
	var data[]enticonfig.MyData
	err = json.Unmarshal(rawMessage[1], &c.Data)
	if err != nil {
		log.Fatal(err)
		return 
	}
	data = append(data, c.Data...)
	s := a.logic
	config := s.Create(ctx, name, data)
	writeJsonResponse(w, http.StatusCreated, config)

}

func(a *api) GetConfig(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	name := chi.URLParam(r, "service")

	conf, _ := a.logic.GetConfig(ctx, name)

	writeJsonResponse(w, http.StatusOK, conf)
}

func(a *api) DeleteConfig(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	name := chi.URLParam(r, "service")
	conf, _ := a.logic.DeleteConfig(ctx, name)

	writeJsonResponse(w, http.StatusOK, conf)

}

func (a *api) Serve() error {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		writeJsonResponse(w, http.StatusAccepted, "hello")
	})
	apiRouter.Post("/config", http.HandlerFunc(a.CreateService))
	a.server = &http.Server{Addr: a.config.ServeAddress, Handler: apiRouter}

	apiRouter.Get("/config", http.HandlerFunc(a.GetConfig))
	a.server = &http.Server{Addr: a.config.ServeAddress, Handler: apiRouter}
	apiRouter.Delete("/config", http.HandlerFunc(a.DeleteConfig))
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
