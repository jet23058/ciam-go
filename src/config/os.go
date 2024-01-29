package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	SERVER_PORT string
	DOMAIN      string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	DB_DATABASE string
	DB_SSLMODE  string
	JWT_SECRET  []byte
	JWT_ISSUER  string
}

var config *envConfig

func GetEnv() *envConfig {
	if config == nil {
		config = &envConfig{
			SERVER_PORT: os.Getenv("SERVER_PORT"),
			DOMAIN:      os.Getenv("DOMAIN"),
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_USER:     os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_PORT:     os.Getenv("DB_PORT"),
			DB_DATABASE: os.Getenv("DB_DATABASE"),
			DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
			JWT_SECRET:  []byte(os.Getenv("JWT_SECRET")),
			JWT_ISSUER:  os.Getenv("JWT_ISSUER"),
		}
	}
	return config
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitOs() {
	initEnv()
	os.Setenv("TZ", "0")
}
