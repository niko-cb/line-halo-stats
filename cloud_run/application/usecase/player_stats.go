package usecase

import (
	"context"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/domain/service"
	"line-halo-stats/cloud_run/lib/log"
)

type IPlayerStatsUsecase interface {
	GetPlayerStats(ctx context.Context, gamertag string) (*entity.PlayerStats, error)
}

var _ IPlayerStatsUsecase = &PlayerStatsUsecase{}

type PlayerStatsUsecase struct {
	hs service.IAutocodeService
}

func NewPlayerStatsUsecase(hs service.IAutocodeService) IPlayerStatsUsecase {
	return &PlayerStatsUsecase{
		hs: hs,
	}
}

func (psu *PlayerStatsUsecase) GetPlayerStats(ctx context.Context, gamertag string) (*entity.PlayerStats, error) {
	playerCSRS, err := psu.hs.GetPlayerCSRS(ctx, gamertag)
	if err != nil {
		log.Errorf(ctx, "%+v", err)
	}

	playerServiceRecord, err := psu.hs.GetPlayerServiceRecord(ctx, gamertag)
	if err != nil {
		log.Errorf(ctx, "%+v", err)
	}

	return psu.hs.GetPlayerStats(ctx, playerCSRS, playerServiceRecord)
}
