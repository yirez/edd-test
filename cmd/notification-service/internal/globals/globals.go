package globals

import (
	"edd-test/cmd/notification-service/internal/types"
	"github.com/yirez/go-gw-test/pkg/configuration_manager"
	cmt "github.com/yirez/go-gw-test/pkg/configuration_manager/types"

	"go.uber.org/zap"
)

var Cfg types.AppConfig

func InitConfiguration() error {
	standardConfigs, err := configuration_manager.InitStandardConfigs(
		cmt.InitChecklist{
			DB: true,
		},
	)
	if err != nil {
		return err
	}

	Cfg.StandardConfigs = standardConfigs
	if err := configuration_manager.ReadCustomConfig("service_name", &Cfg.ServiceName); err != nil {
		return err
	}

	zap.ReplaceGlobals(Cfg.StandardConfigs.Clients.Logger)
	return nil
}
