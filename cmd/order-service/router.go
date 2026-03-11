package main

import (
	"edd-test/cmd/order-service/internal/repo"
	"net/http"

	"edd-test/cmd/order-service/internal/globals"
	"edd-test/cmd/order-service/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/yirez/go-gw-test/pkg/rest_qol"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	orderUseCase := usecase.NewOrderUseCase(repo.NewOrderRepo())

	router.Use(rest_qol.RequestIDMiddleware(globals.Cfg.ServiceName))
	router.Use(rest_qol.AccessLoggingMiddleware())
	rest_qol.RegisterOperationalRoutes(router, nil, nil)
	router.HandleFunc("/orders", orderUseCase.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}", orderUseCase.GetOrder).Methods(http.MethodGet)
	router.HandleFunc("/orders/{id}/cancel", orderUseCase.CancelOrder).Methods(http.MethodPost)

	return router
}
