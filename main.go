package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/subkhiyoga/app-mahasiswa-api/controller"
	"github.com/subkhiyoga/app-mahasiswa-api/repository"
	"github.com/subkhiyoga/app-mahasiswa-api/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// config db
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "uhuy123"
	dbName := "testdb"
	sslMode := "disable"
	serverPort := ":8080"

	// db connection
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// gin router

	router := gin.Default()

	// group users
	mahasiswaRouter := router.Group("/api/v1/mahasiswa")

	// dependencies (repository, usecase, controller)
	mahasiswaRepo := repository.NewMahasiswaRepo(db)
	mahasiswaUsecase := usecase.NewMahasiswaUsecase(mahasiswaRepo)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaUsecase)

	// routes (GET, POST, PUT, DELETE)
	mahasiswaRouter.GET("", mahasiswaController.FindData)
	mahasiswaRouter.GET("/:id", mahasiswaController.FindDataById)
	mahasiswaRouter.POST("", mahasiswaController.Register)
	mahasiswaRouter.PUT("", mahasiswaController.Edit)
	mahasiswaRouter.DELETE("/:id", mahasiswaController.Unreg)

	// run server
	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}
