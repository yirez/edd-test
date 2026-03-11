package main

import (
	"fmt"

	"edd-test/cmd/order-service/internal/globals"
	"github.com/yirez/go-gw-test/pkg/rest_qol"

	"go.uber.org/zap"
)

func main() {
	err := globals.InitConfiguration()
	if err != nil {
		fmt.Printf("failed init configs: %v\n", err)
		return
	}

	router := NewRouter()
	address := fmt.Sprintf(":%d", globals.Cfg.StandardConfigs.Port)
	err = rest_qol.RunHTTPServer(address, router)
	if err != nil {
		zap.L().Error("server shutdown", zap.Error(err))
	}
}
