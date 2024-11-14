package eods_controller

import (
	"gin-gorm/app/model"
	"gin-gorm/app/request"
	"gin-gorm/app/response"
	"gin-gorm/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

const TABLE = "eods"

func GetAllData(ctx *gin.Context) {
	data_list := new([]model.Eod)

	err := database.DB.Table(TABLE).Find(&data_list).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data_list,
		"name": ctx.MustGet("name").(string),
	})
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	data := new(response.EodResponse)

	err := database.DB.Table(TABLE).Where("id=?", id).Find(&data).Error

	if err != nil || data.ID == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func Store(ctx *gin.Context) {
	dataReq := new(request.EodRequest)

	if errReq := ctx.ShouldBind(dataReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	store := new(model.Eod)

	store.Name = &dataReq.Name

	errDb := database.DB.Table(TABLE).Create(&store).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "fail to create data",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
	})
}

func Update(ctx *gin.Context) {
	dataReq := new(request.EodRequest)

	if errReq := ctx.ShouldBind(dataReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	id := ctx.Param("id")
	data := new(model.Eod)
	errDb := database.DB.Table(TABLE).Where("id=?", id).Find(&data).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "fail to update data",
		})
		return
	}

	update := new(model.Eod)

	update.Name = &dataReq.Name

	errUpdate := database.DB.Table(TABLE).Where("id=?", id).Updates(&update).Error

	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "fail to update data",
		})
		return
	}
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	data := new(model.Eod)
	errDb := database.DB.Table(TABLE).Where("id=?", id).Find(&data).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "fail to delete data",
		})
		return
	}

	errDelete := database.DB.Table(TABLE).Unscoped().Where("id=?", id).Delete(&model.Eod{}).Error

	if errDelete != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "fail to update data",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}
