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
	redisPassword := getEnv("REDIS_PASS", "isucon")
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       int(redisDB),
	})
}

const scoreKey = "UserRanking"

func addScoreToUser(userID string, score int) error {
	return nil
}

func addScoreToUserImpl(ctx context.Context, userID int64, score int) error {
	return redisClient.ZAdd(ctx, scoreKey, &redis.Z{Score: float64(score), Member: userID}).Err()
}

func getRankOfUser(userID string) (int, error) {
	return 0, nil
}

func getRankOfUserImpl(ctx context.Context, userID int64) (int, error) {
	rank, err := redisClient.ZRank(ctx, scoreKey, strconv.FormatInt(userID, 10)).Result()
	if err != nil {
		return 0, err
	}
	return int(rank), err
}
