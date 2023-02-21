package controllers

import (
	"net/http"

	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"
)

type MahasiswaController interface {
	CreateMahasiswa()(ctx *gin.Context)
	ReadMahasiswa()(ctx *gin.Context)
	UpdateMahasiswa()(ctx *gin.Context)
	DeleteMahasiswa()(ctx *gin.Context)
}

type mahasiswaController struct {
	mahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaService service.MahasiswaService) *mahasiswaController {
	return &mahasiswaController{mahasiswaService}
}

func (c *mahasiswaController) CreateMahasiswa(ctx *gin.Context) {
	var mahasiswa entity.Mahasiswa

	err := ctx.ShouldBindJSON(&mahasiswa)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot proccess entity",
		})
		return
	}

	newMaha, err := c.mahasiswaService.CreateMahasiswa(mahasiswa)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot make new Mahasiswa",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "insert Mahasiswa success",
		"result":  newMaha,
	})
}

func (c *mahasiswaController) ReadMahasiswa(ctx *gin.Context) {
	var result gin.H
	dos, err := c.mahasiswaService.ReadMahasiswa()
	if err != nil {
		result = gin.H{
			"message": err,
		}
		return
	} else {
		result = gin.H{
			"result": dos,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *mahasiswaController) UpdateMahasiswa(ctx *gin.Context)  {
	var mahasiswa entity.Mahasiswa

	nim := ctx.Param("nim")

	err := ctx.ShouldBindJSON(&mahasiswa)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot bind to json",
		})
		return
	}

	newMaha, err := c.mahasiswaService.UpdateMahasiswa(mahasiswa, nim)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"result":  newMaha,
		})		
	}else{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot update Mahasiswas",
		})
		return
	}
}

func (c *mahasiswaController) DeleteMahasiswa(ctx *gin.Context) {

	nim := ctx.Param("nim")


	err := c.mahasiswaService.DeleteMahasiswa(nim)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot delete category",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
	})
}