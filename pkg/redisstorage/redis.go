package redisstorage

import (
	"runtime"
	"sudoku/config"

	"github.com/gofiber/storage/redis/v2"
)

func NewRedisStorage(config config.RedisConfig) *redis.Storage {
	redisStore := redis.New(redis.Config{
		Host:      config.Url,
		Port:      config.Port,
		Database:  config.DB,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	return redisStore
}
