package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"sync/atomic"
	"task/config"
	"task/internal/service"
	"time"
)

type redisRepository struct {
	cfg *config.Config
	db  *redis.Client
	id  atomic.Uint64
}

func NewRedisRepository(cfg *config.Config, db *redis.Client) service.Repository {
	return &redisRepository{db: db, cfg: cfg}
}

func (r *redisRepository) CntRequest(ctx context.Context, userID int64) (bool, error) {
	i, err := r.db.ZRemRangeByScore(
		ctx,
		strconv.FormatInt(r.cfg.Server.Number, 10)+strconv.FormatInt(userID, 10),
		"-inf",
		strconv.FormatInt(time.Now().Unix(), 10)).Result()
	log.Println("ZREM:::", i)
	if err != nil {
		log.Println("ZREMERR:::", err)
		return false, err
	}

	result, err := r.db.ZCount(
		ctx,
		strconv.FormatInt(r.cfg.Server.Number, 10)+strconv.FormatInt(userID, 10),
		"-inf",
		"+inf").Result()
	log.Println("ZCount:::", result)
	if result > r.cfg.Server.K-1 {
		log.Println("ZCountERR:::", err)
		return false, fmt.Errorf("flood")
	}

	r.id.Add(1)
	zadd, err := r.db.ZAdd(
		ctx,
		strconv.FormatInt(r.cfg.Server.Number, 10)+strconv.FormatInt(userID, 10),
		redis.Z{
			Score:  float64(time.Now().Add(r.cfg.Server.N).Unix()),
			Member: r.id.Load(),
		},
	).Result()
	log.Println("ZADD:::", zadd)
	if err != nil {
		log.Println("ZADDERR:::", zadd)
		return false, err
	}

	return true, nil
}
