package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

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

const userIconHashKeyPrefix = "UserIconHash:"

func userIconHashKey(userName string) string {
	return fmt.Sprintf("%s:%s", userIconHashKeyPrefix, userName)
}

func setUserIconSha256(ctx context.Context, userName string, sha256 string) error {
	return redisClient.Set(ctx, userIconHashKey(userName), sha256, 1*time.Minute).Err()
}

func getUserIconSha256(ctx context.Context, userName string) (string, error) {
	return redisClient.Get(ctx, userIconHashKey(userName)).Result()
}
