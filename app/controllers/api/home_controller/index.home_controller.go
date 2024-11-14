package home_controller

import (
	"gin-gorm/app/model"
	"gin-gorm/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SplashScreen(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": true,
		"text":   "success",
		"result": gin.H{
			"latest_version": 2,
			"time":           time.Now().Format("15:04:05"),
		},
	})
}

func VersionCheck(ctx *gin.Context) {
	app_version := ctx.DefaultQuery("app_version", "1")

	app_version_int, _ := strconv.Atoi(app_version)

	var data_karyawan model.Karyawan
	err_karyawan := database.DB.Table("karyawan").
		Where("karyawan_email = ?", "tries1@mail.com").
		First(&data_karyawan).Error

	if err_karyawan != nil {
		ctx.JSON(202, gin.H{
			"status": false,
			"text":   "ERR DB karyawan",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": true,
		"text":   "success",
		"result": gin.H{
			"app_version": app_version_int,
			"api_version": 3,
		},
	})
}
