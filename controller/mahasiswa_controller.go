package controller

import (
	"net/http"
	"strconv"

	"github.com/subkhiyoga/app-mahasiswa-api/model"
	"github.com/subkhiyoga/app-mahasiswa-api/model/response"
	"github.com/subkhiyoga/app-mahasiswa-api/usecase"

	"github.com/gin-gonic/gin"
)

type MahasiswaController struct {
	usecase usecase.MahasiswaUsecase
}

func (c *MahasiswaController) FindData(ctx *gin.Context) {
	result := c.usecase.FindData()
	if result == nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusInternalServerError, "Failed to get students")
		return
	}
	response.JSONSuccess(ctx.Writer, http.StatusOK, result)
}

func (c *MahasiswaController) FindDataById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid user ID")
		return
	}

	result := c.usecase.FindDataById(id)
	response.JSONSuccess(ctx.Writer, http.StatusOK, result)
}

func (c *MahasiswaController) Register(ctx *gin.Context) {
	var newMahasiswa model.Mahasiswa

	if err := ctx.BindJSON(&newMahasiswa); err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid Input")
		return
	}

	result := c.usecase.Register(&newMahasiswa)
	response.JSONSuccess(ctx.Writer, http.StatusCreated, result)
}

func (c *MahasiswaController) Edit(ctx *gin.Context) {
	var data model.Mahasiswa

	if err := ctx.BindJSON(&data); err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&data)
	response.JSONSuccess(ctx.Writer, http.StatusOK, res)
}

func (c *MahasiswaController) Unreg(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid user ID")
		return
	}

	result := c.usecase.Unreg(id)
	response.JSONSuccess(ctx.Writer, http.StatusOK, result)
}

func NewMahasiswaController(u usecase.MahasiswaUsecase) *MahasiswaController {
	controller := MahasiswaController{
		usecase: u,
	}

	return &controller
}
