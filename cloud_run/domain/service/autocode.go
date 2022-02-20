package service

import (
	"context"
	"line-halo-stats/cloud_run/domain/entity"
	"line-halo-stats/cloud_run/domain/repository"
)

type IAutocodeService interface {
	GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error)
	GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.ServiceRecord, error)
	GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.ServiceRecord) (*entity.PlayerStats, error)
}

type AutocodeService struct {
	repo repository.AutocodeRepository
}

func NewAutocodeService(repo repository.AutocodeRepository) IAutocodeService {
	return &AutocodeService{
		repo: repo,
	}
}

func (s AutocodeService) GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error) {
	return s.repo.GetPlayerCSRS(ctx, gamertag)
}

func (s AutocodeService) GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.ServiceRecord, error) {
	return s.repo.GetPlayerServiceRecord(ctx, gamertag)
}

func (s AutocodeService) GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.ServiceRecord) (*entity.PlayerStats, error) {
	return s.repo.GetPlayerStats(ctx, csrs, sr)
}
