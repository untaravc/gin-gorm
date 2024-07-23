package karyawan_controller

import (
	"gin-gorm/database"
	"gin-gorm/models"

	"github.com/gin-gonic/gin"
)

const TABLE = "eods"

func GetAllUser(ctx *gin.Context) {
	karyawan := new([]models.Karyawan)

	err := database.DB.Table(TABLE).Find(&karyawan).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": karyawan,
	})
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	karyawan := new(models.Karyawan)

	err := database.DB.Table(TABLE).Where("id=?", id).Find(&karyawan).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": karyawan,
	})
}
