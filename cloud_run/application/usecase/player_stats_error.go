package usecase

import (
	"context"
	"line-halo-stats/cloud_run/domain/errors"
)

type IPlayerStatsErrorUsecase interface {
	ErrorHandle(ctx context.Context, e *errors.PlayerStatsRetrievalError) error
}

var _ IPlayerStatsErrorUsecase = &PlayerStatsErrorUsecase{}

type PlayerStatsErrorUsecase struct {
}

func NewPlayerStatsErrorUsecase() IPlayerStatsErrorUsecase {
	return &PlayerStatsErrorUsecase{}
}

func (u *PlayerStatsErrorUsecase) ErrorHandle(ctx context.Context, e *errors.PlayerStatsRetrievalError) error {
	return e.Err
}
