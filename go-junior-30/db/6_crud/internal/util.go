package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DRIVER string
	DB_URL    string
	HOST      string
}

func NewConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB_DRIVER: os.Getenv("DB_DRIVER"),
		DB_URL:    os.Getenv("DB_URL"),
		HOST:      os.Getenv("HOST"),
	}, nil
}
