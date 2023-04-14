package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string) string {
	// load .env file
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("error saat load .env file")
	}

	return os.Getenv(key)
}
