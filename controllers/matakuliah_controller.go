package controllers

import (
	"net/http"

	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"
)

type MataKuliahController interface {
	CreateMataKuliah()(ctx *gin.Context)
	ReadMataKuliah()(ctx *gin.Context)
	UpdateMataKuliah()(ctx *gin.Context)
	DeleteMataKuliah()(ctx *gin.Context)

}

type mataKuliahController struct {
	mataKuliahService service.MataKuliahService
}

func NewMataKuliahController(mataKuliahService service.MataKuliahService) *mataKuliahController {
	return &mataKuliahController{mataKuliahService}
}

func (c *mataKuliahController) CreateMataKuliah(ctx *gin.Context) {
	var mKul entity.MataKuliah

	err := ctx.ShouldBindJSON(&mKul)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot proccess entity",
		})
		return
	}

	newMKul, err := c.mataKuliahService.CreateMataKuliah(mKul)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot make new MataKuliah",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "insert MataKuliah success",
		"result":  newMKul,
	})
}

func (c *mataKuliahController) ReadMataKuliah(ctx *gin.Context) {
	var result gin.H
	mKul, err := c.mataKuliahService.ReadMataKuliah()
	if err != nil {
		result = gin.H{
			"message": err,
		}
		return
	} else {
		result = gin.H{
			"result": mKul,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *mataKuliahController) UpdateMataKuliah(ctx *gin.Context)  {
	var mKul entity.MataKuliah

	id := ctx.Param("id")

	err := ctx.ShouldBindJSON(&mKul)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot bind to json",
		})
		return
	}

	newMKul, err := c.mataKuliahService.UpdateMataKuliah(mKul, id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"result":  newMKul,
		})		
	}else{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot update Matakuliah",
		})
		return
	}
}

func (c *mataKuliahController)	DeleteMataKuliah(ctx *gin.Context){

	id := ctx.Param("id")

	err := c.mataKuliahService.DeleteMataKuliah(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot delete Matakuliah",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Matakuliah",
	})
}