package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DataBaseConfig struct {
	Url string
}

type LogConfig struct {
	Level  int
	Format string
}

type RedisConfig struct {
	Port        int           `yaml:"port"`
	Url         string        `yaml:"url"`
	Password    string        `yaml:"password"`
	User        string        `yaml:"user"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Port:        getInt("REDIS_PORT", 6379),
		Url:         getString("REDIS_URL", "127.0.0.1"),
		Password:    getString("REDIS_PASS", "my_pass"),
		User:        getString("REDIS_USER", "user"),
		DB:          getInt("REDIS_DATABASE", 0),
		MaxRetries:  getInt("REDIS_MAXRET", 5),
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}
	log.Println(".env file loaded")
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}
}

func NewDBConfig() *DataBaseConfig {
	databaseURL := os.Getenv("DB_URL")
	return &DataBaseConfig{
		Url: databaseURL,
	}
}

func getString(key, defValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defValue
	}
	return value
}

func getInt(key string, defValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defValue
	}
	if value == "" {
		return defValue
	}
	return i
}
