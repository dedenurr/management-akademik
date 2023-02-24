package controllers

import (
	"net/http"
	"strconv"

	"github.com/dedenurr/management-akademik/entity"
	"github.com/dedenurr/management-akademik/service"
	"github.com/gin-gonic/gin"
)

type PerkuliahanController interface {
	CreatePerkuliahan()(ctx *gin.Context)
	ReadPerkuliahan()(ctx *gin.Context)
	UpdatePerkuliahan()(ctx *gin.Context)
	DeletePerkuliahan()(ctx *gin.Context)
}

type perkuliahanController struct {
	perkuliahanService service.PerkuliahanService
}

func NewPerkuliahanController(perkuliahanService service.PerkuliahanService) *perkuliahanController {
	return &perkuliahanController{perkuliahanService}
}

func (c *perkuliahanController) CreatePerkuliahan(ctx *gin.Context) {
	var inputPerkuliahan entity.InputPerkuliahan

	err := ctx.ShouldBindJSON(&inputPerkuliahan)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot proccess entity",
		})
		return
	}

	newPkh, err := c.perkuliahanService.CreatePerkuliahan(inputPerkuliahan)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot make new Perkuliahan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "insert Perkuliahan success",
		"result":  newPkh,
	})
}

func (c *perkuliahanController) ReadPerkuliahan(ctx *gin.Context) {
	var result gin.H
	pkh, err := c.perkuliahanService.ReadPerkuliahan()
	if err != nil {
		result = gin.H{
			"message": err,
		}
		return
	} else {
		result = gin.H{
			"result": pkh,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *perkuliahanController) UpdatePerkuliahan(ctx *gin.Context)  {
	var pkh entity.InputPerkuliahan

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "cannot parse id into integer",
		})
		return
	}


	err = ctx.ShouldBindJSON(&pkh)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot bind to json",
		})
		return
	}

	newPkh, err := c.perkuliahanService.UpdatePerkuliahan(pkh, id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"result":  newPkh,
		})		
	}else{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot update Perkuliahans",
		})
		return
	}
}

func (c *perkuliahanController) DeletePerkuliahan(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "cannot parse id into integer",
		})
		return
	}

	err = c.perkuliahanService.DeletePerkuliahan(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot delete Perkuliahan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Perkuliahan",
	})
}