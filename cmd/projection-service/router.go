package main

import (
	"edd-test/cmd/projection-service/internal/repo"
	"net/http"

	"edd-test/cmd/projection-service/internal/globals"
	"edd-test/cmd/projection-service/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/yirez/go-gw-test/pkg/rest_qol"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	projectionUseCase := usecase.NewProjectionUseCase(repo.NewOrderProjectionRepo())

	router.Use(rest_qol.RequestIDMiddleware(globals.Cfg.ServiceName))
	router.Use(rest_qol.AccessLoggingMiddleware())
	rest_qol.RegisterOperationalRoutes(router, nil, nil)
	router.HandleFunc("/projections/orders/{id}", projectionUseCase.GetOrderView).Methods(http.MethodGet)
	router.HandleFunc("/workers/events", projectionUseCase.ApplyEvent).Methods(http.MethodPost)

	return router
}
