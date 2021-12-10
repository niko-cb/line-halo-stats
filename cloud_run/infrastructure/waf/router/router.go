package router

import (
	"line-halo-stats/cloud_run/di"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/infrastructure/waf/config"
	"line-halo-stats/cloud_run/infrastructure/waf/handler"
)

func InitializePlayerStatsHandler(cfg *config.ServerConfig) handler.IPlayerStatsHandler {
	return di.InitializePlayerStatsHandler(cfg, entity.HalodotapiURL(cfg.HalodotapiURL))
}
