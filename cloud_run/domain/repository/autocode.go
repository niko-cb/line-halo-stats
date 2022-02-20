package repository

import (
	"context"
	"line-halo-stats/cloud_run/domain/entity"
)

type AutocodeRepository interface {
	GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error)
	GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.ServiceRecord, error)
	GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.ServiceRecord) (*entity.PlayerStats, error)
}
