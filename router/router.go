package router

import (
	"log"

	"github.com/subkhiyoga/app-mahasiswa-api/config"
	"github.com/subkhiyoga/app-mahasiswa-api/controller"
	"github.com/subkhiyoga/app-mahasiswa-api/repository"
	"github.com/subkhiyoga/app-mahasiswa-api/usecase"
	"github.com/subkhiyoga/app-mahasiswa-api/utils"

	"github.com/subkhiyoga/auth-jwt/controll"
	"github.com/subkhiyoga/auth-jwt/repo"
	"github.com/subkhiyoga/auth-jwt/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func StartServer() {
	s_key := []byte(utils.DotEnv("test"))

	db := config.LoadDB()
	defer db.Close()
	authMiddleware := controll.AuthMiddleware(s_key)

	router := gin.Default()
	mahasiswaRouter := router.Group("/api/v1/mahasiswa")
	mahasiswaRouter.Use(authMiddleware)

	mrepository := repository.NewMahasiswaRepo(db)
	musecase := usecase.NewMahasiswaUsecase(mrepository)
	mcontroller := controller.NewMahasiswaController(musecase)
	mrepo := repo.NewMahasiswaRepo(db)
	loginService := service.NewLoginService(mrepo)
	loginJwt := controll.NewCredentialsJwt(loginService)

	router.POST("/auth/login", loginJwt.Login)
	router.POST("/register", mcontroller.Register)
	mahasiswaRouter.GET("", mcontroller.FindData)
	mahasiswaRouter.GET("/:id", mcontroller.FindDataById)
	mahasiswaRouter.PUT("", mcontroller.Edit)
	mahasiswaRouter.DELETE("/:id", mcontroller.Unreg)

	// run server
	err := router.Run(utils.DotEnv("SERVER_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

}
