package config

import (
	"github.com/joho/godotenv"
	"gitlab.com/nevasik7/lg"
	"os"
	"strconv"
)

type PostgresConfig struct {
	DSN         string
	SQLQuery    string
	DurationMs  int
	Concurrency int
}

func MustLoad() (PostgresConfig, error) {
	lg.Init()
	err := godotenv.Load(".env")
	if err != nil {
		lg.Fatalf("Не удалось загрузить .env файл: %v", err)
	}

	return PostgresConfig{
		DSN:         os.Getenv("DSN"),
		SQLQuery:    os.Getenv("SQL_QUERY"),
		DurationMs:  getEnvInt("DURATION_MS", 5000),
		Concurrency: getEnvInt("CONCURRENCY", 10),
	}, nil

}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		return atoi(value, defaultValue)
	}
	return defaultValue
}

func atoi(str string, defaultValue int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}
