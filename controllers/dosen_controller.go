package controllers

import (
	"net/http"

	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"
)

type DosenController interface {
	CreateDosen()(ctx *gin.Context)
	ReadDosen()(ctx *gin.Context)
	UpdateDosen()(ctx *gin.Context)
	DeleteDosen()(ctx *gin.Context)
}

type dosenController struct {
	dosenService service.DosenService
}

func NewDosenController(dosenService service.DosenService) *dosenController {
	return &dosenController{dosenService}
}

func (c *dosenController) CreateDosen(ctx *gin.Context) {
	var dosen entity.Dosen

	err := ctx.ShouldBindJSON(&dosen)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot proccess entity",
		})
		return
	}

	newDos, err := c.dosenService.CreateDosen(dosen)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot make new dosen",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "insert dosen success",
		"result":  newDos,
	})
}

func (c *dosenController) ReadDosen(ctx *gin.Context) {
	var result gin.H
	dos, err := c.dosenService.ReadDosen()
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

func (c *dosenController) UpdateDosen(ctx *gin.Context)  {
	var dos entity.Dosen

	nip := ctx.Param("nip")

	err := ctx.ShouldBindJSON(&dos)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot bind to json",
		})
		return
	}

	newDos, err := c.dosenService.UpdateDosen(dos, nip)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"result":  newDos,
		})		
	}else{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot update Dosens",
		})
		return
	}
}

func (c *dosenController) DeleteDosen(ctx *gin.Context) {

	nip := ctx.Param("nip")


	err := c.dosenService.DeleteDosen(nip)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot delete Dosen",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Dosen",
	})
}