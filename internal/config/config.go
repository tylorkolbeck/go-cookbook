package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST      string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
	DB_PORT      string
	PORT         string
	TOKEN_SECRET []byte
}

func NewConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, errors.New("Error loading .env file")
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return Config{}, errors.New("DB_HOST not set")
	}
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return Config{}, errors.New("DB_USER not set")
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return Config{}, errors.New("DB_PASSWORD not set")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return Config{}, errors.New("DB_NAME not set")
	}
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return Config{}, errors.New("DB_PORT not set")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return Config{}, errors.New("PORT not set")
	}
	tokenSecret, ok := os.LookupEnv("TOKEN_SECRET")
	if !ok {
		return Config{}, errors.New("TOKEN_SECRET not set")
	}

	return Config{
		DB_HOST:      dbHost,
		DB_USER:      dbUser,
		DB_PASSWORD:  dbPassword,
		DB_NAME:      dbName,
		DB_PORT:      dbPort,
		PORT:         port,
		TOKEN_SECRET: []byte(tokenSecret),
	}, nil

}
