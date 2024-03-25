package usecase

import (
	"context"
	"task/config"
	"task/internal/service"
)

type FloodControlUsecase struct {
	cfg  *config.Config
	repo service.Repository
}

func NewServiceUsecase(cfg *config.Config, repo service.Repository) service.FloodControl {
	return &FloodControlUsecase{cfg: cfg, repo: repo}
}

func (u *FloodControlUsecase) Check(ctx context.Context, userID int64) (bool, error) {
	return u.repo.CntRequest(ctx, userID)
}
