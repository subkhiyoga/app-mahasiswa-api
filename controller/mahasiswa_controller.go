package controller

import (
	"net/http"
	"strconv"

	"app-mahasiswa-api/model"
	"app-mahasiswa-api/usecase"

	"github.com/gin-gonic/gin"
)

type MahasiswaController struct {
	usecase usecase.MahasiswaUsecase
}

func (c *MahasiswaController) FindData(ctx *gin.Context) {
	result := c.usecase.FindData()

	ctx.JSON(http.StatusOK, result)
}

func (c *MahasiswaController) FindDataById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data ID")
		return
	}

	result := c.usecase.FindDataById(id)
	ctx.JSON(http.StatusOK, result)
}

func (c *MahasiswaController) Register(ctx *gin.Context) {
	var newData model.Mahasiswa

	err := ctx.BindJSON(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	result := c.usecase.Register(&newData)
	ctx.JSON(http.StatusCreated, result)
}

func (c *MahasiswaController) Edit(ctx *gin.Context) {
	var data model.Mahasiswa

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
	}

	result := c.usecase.Edit(&data)
	ctx.JSON(http.StatusOK, result)
}

func (c *MahasiswaController) Unreg(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data ID")
		return
	}

	result := c.usecase.Unreg(id)
	ctx.JSON(http.StatusOK, result)
}

func NewMahasiswaController(u usecase.MahasiswaUsecase) *MahasiswaController {
	controller := MahasiswaController{
		usecase: u,
	}

	return &controller
}
