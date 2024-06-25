package appconfig

import (
	"log"
	"webhooklenna/config"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.DB = config.Postgres()
}
