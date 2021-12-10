package controller

import (
	"context"
	"line-halo-stats/cloud_run/application/usecase"
	"line-halo-stats/cloud_run/domain/errors"
)

type ErrorController struct {
	use usecase.IPlayerStatsErrorUsecase
}

func NewErrorController(use usecase.IPlayerStatsErrorUsecase) *ErrorController {
	return &ErrorController{use: use}
}

func (con *ErrorController) HandleError(ctx context.Context, e *errors.PlayerStatsRetrievalError) error {
	return con.use.ErrorHandle(ctx, e)
}
