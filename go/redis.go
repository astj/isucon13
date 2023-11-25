package main

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

func getEnv(key string, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func connectRedis(logger echo.Logger) *redis.Client {
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	return redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   int(redisDB),
	})
}

const scoreKey = "UserRanking"

func addScoreToUser(ctx context.Context, userName string, score int) error {
	return redisClient.ZIncrBy(ctx, scoreKey, float64(score), userName).Err()
}

func getRankOfUser(ctx context.Context, userName string) (int, error) {
	rank, err := redisClient.ZRevRank(ctx, scoreKey, userName).Result()
	if err != nil {
		return 0, err
	}
	// revrank は 0 が最小なので +1 する
	return int(rank) + 1, err
}
