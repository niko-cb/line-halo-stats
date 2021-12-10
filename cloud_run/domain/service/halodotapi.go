package service

import (
	"context"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/domain/repository"
)

type IHalodotapiService interface {
	GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error)
	GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.PlayerServiceRecord, error)
	GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.PlayerServiceRecord) (*entity.PlayerStats, error)
}

type HalodotapiService struct {
	repo repository.HalodotapiRepository
}

func NewHalodotapiService(repo repository.HalodotapiRepository) IHalodotapiService {
	return &HalodotapiService{
		repo: repo,
	}
}

func (s HalodotapiService) GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error) {
	return s.repo.GetPlayerCSRS(ctx, gamertag)
}

func (s HalodotapiService) GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.PlayerServiceRecord, error) {
	return s.repo.GetPlayerServiceRecord(ctx, gamertag)
}

func (s HalodotapiService) GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.PlayerServiceRecord) (*entity.PlayerStats, error) {
	return s.repo.GetPlayerStats(ctx, csrs, sr)
}
