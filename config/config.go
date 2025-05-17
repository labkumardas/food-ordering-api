package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ApiKey string

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ApiKey = os.Getenv("API_KEY")
	if ApiKey == "" {
		log.Fatal("API_KEY not found in environment variables")
	}
}
