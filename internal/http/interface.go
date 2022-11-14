package http

import (
	"context"
	"net/http"
)

type API interface{
	Serve() error
	Shutdown(ctx context.Context)

	CreateService(w http.ResponseWriter, r *http.Request)
	GetConfig(w http.ResponseWriter, r *http.Request)
}