package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvDatabase struct {
	Host     string
	Database string
	Password string
	User     string
}

type EnvAuth struct {
	Username string
	Password string
}

type EnvConfig struct {
	Port     string
	Database EnvDatabase
	Auth     EnvAuth
}

func NewConfig() EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return EnvConfig{
		Port: os.Getenv("PORT"),
		Auth: EnvAuth{
			Username: os.Getenv("AUTH_USER"),
			Password: os.Getenv("AUTH_PASS"),
		},
		Database: EnvDatabase{
			Host:     os.Getenv("DB_HOST"),
			Database: os.Getenv("DB_DATABASE"),
			Password: os.Getenv("DB_PASSWORD"),
			User:     os.Getenv("DB_USER"),
		},
	}
}
