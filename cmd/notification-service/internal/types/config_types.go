package types

import cmt "github.com/yirez/go-gw-test/pkg/configuration_manager/types"

type AppConfig struct {
	StandardConfigs cmt.StandardConfig
	ServiceName     string
}
