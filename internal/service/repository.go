package service

import "context"

type Repository interface {
	CntRequest(ctx context.Context, userID int64) (bool, error)
}
