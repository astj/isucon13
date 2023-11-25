package main

import (
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
