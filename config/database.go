package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/subkhiyoga/app-mahasiswa-api/utils"
)

func LoadDB() *sql.DB {
	dbHost := utils.DotEnv("DB_HOST")
	dbPort := utils.DotEnv("DB_PORT")
	dbUser := utils.DotEnv("DB_USER")
	dbPassword := utils.DotEnv("DB_PASSWORD")
	dbName := utils.DotEnv("DB_NAME")
	sslMode := utils.DotEnv("SSL_MODE")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} else {
		log.Println("Database Successfully Connected")
	}

	return db
}
