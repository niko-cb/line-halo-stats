// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"line-halo-stats/cloud_run/application/usecase"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/domain/service"
	"line-halo-stats/cloud_run/infrastructure/gateway/halodotapi"
	"line-halo-stats/cloud_run/infrastructure/waf/config"
	"line-halo-stats/cloud_run/infrastructure/waf/handler"
	"line-halo-stats/cloud_run/interface/controller"
)

// Injectors from wire.go:

// InitializeOrderNumberHandler DI用のイニシャライザー
// salesforceURLを引数に持っているのはcommon_moduleにcfgを直接設定を渡して、サーバー設定に依存させたくないため
func InitializePlayerStatsHandler(cfg *config.ServerConfig, halotrackerURL entity.HalodotapiURL) handler.IPlayerStatsHandler {
	halodotapiRepository := halodotapi.NewHalodotapiRepository(cfg)
	iHalodotapiService := service.NewHalodotapiService(halodotapiRepository)
	iPlayerStatsUsecase := usecase.NewPlayerStatsUsecase(iHalodotapiService)
	playerStatsController := controller.NewPlayerStatsController(iPlayerStatsUsecase)
	iPlayerStatsHandler := handler.NewPlayerStatsHandler(playerStatsController)
	return iPlayerStatsHandler
}