package repository

import (
	"context"
	"line-halo-stats/cloud_run/domain/entity"
)

type HalodotapiRepository interface {
	GetPlayerCSRS(ctx context.Context, gamertag string) (*entity.PlayerCSRS, error)
	GetPlayerServiceRecord(ctx context.Context, gamertag string) (*entity.PlayerServiceRecord, error)
	GetPlayerStats(ctx context.Context, csrs *entity.PlayerCSRS, sr *entity.PlayerServiceRecord) (*entity.PlayerStats, error)
}
