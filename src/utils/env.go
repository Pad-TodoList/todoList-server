package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetGoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return os.Getenv(key)
}
