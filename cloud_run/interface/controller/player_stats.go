package controller

import (
	"context"
	"line-halo-stats/cloud_run/application/usecase"
	"line-halo-stats/cloud_run/domain/entity"
)

type PlayerStatsController struct {
	use usecase.IPlayerStatsUsecase
}

func NewPlayerStatsController(use usecase.IPlayerStatsUsecase) *PlayerStatsController {
	return &PlayerStatsController{use: use}
}

func (con *PlayerStatsController) GetPlayerStats(ctx context.Context, gamertag string) (*entity.PlayerStats, error) {
	return con.use.GetPlayerStats(ctx, gamertag)
}
