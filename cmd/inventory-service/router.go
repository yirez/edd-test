package main

import (
	"edd-test/cmd/inventory-service/internal/repo"
	"net/http"

	"edd-test/cmd/inventory-service/internal/globals"
	"edd-test/cmd/inventory-service/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/yirez/go-gw-test/pkg/rest_qol"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	inventoryUseCase := usecase.NewInventoryUseCase(repo.NewInventoryRepo())

	router.Use(rest_qol.RequestIDMiddleware(globals.Cfg.ServiceName))
	router.Use(rest_qol.AccessLoggingMiddleware())
	rest_qol.RegisterOperationalRoutes(router, nil, nil)
	router.HandleFunc("/workers/order-created", inventoryUseCase.HandleOrderCreated).Methods(http.MethodPost)

	return router
}
