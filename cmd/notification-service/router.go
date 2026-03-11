package main

import (
	"edd-test/cmd/notification-service/internal/repo"
	"net/http"

	"edd-test/cmd/notification-service/internal/globals"
	"edd-test/cmd/notification-service/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/yirez/go-gw-test/pkg/rest_qol"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	notificationUseCase := usecase.NewNotificationUseCase(repo.NewDeliveryRepo())

	router.Use(rest_qol.RequestIDMiddleware(globals.Cfg.ServiceName))
	router.Use(rest_qol.AccessLoggingMiddleware())
	rest_qol.RegisterOperationalRoutes(router, nil, nil)
	router.HandleFunc("/workers/events", notificationUseCase.HandleEvent).Methods(http.MethodPost)

	return router
}
