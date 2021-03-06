//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"line-halo-stats/cloud_run/application/usecase"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/domain/service"
	"line-halo-stats/cloud_run/infrastructure/gateway/autocode"
	"line-halo-stats/cloud_run/infrastructure/waf/config"
	"line-halo-stats/cloud_run/infrastructure/waf/handler"
	"line-halo-stats/cloud_run/interface/controller"
)

// InitializeOrderNumberHandler DI用のイニシャライザー
// salesforceURLを引数に持っているのはcommon_moduleにcfgを直接設定を渡して、サーバー設定に依存させたくないため
func InitializePlayerStatsHandler(cfg *config.ServerConfig, a entity.AutocodeURL) handler.IPlayerStatsHandler {
	wire.Build(
		handler.NewPlayerStatsHandler,
		controller.NewPlayerStatsController,
		autocode.NewAutocodeRepository,
		service.NewAutocodeService,
		usecase.NewPlayerStatsUsecase,
		//usecase.NewPlayerStatsErrorUsecase,
		//controller.NewErrorController,
	)
	return &handler.PlayerStatsHandler{}
}
