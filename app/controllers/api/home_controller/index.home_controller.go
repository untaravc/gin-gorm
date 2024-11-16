package home_controller

import (
	"gin-gorm/app/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SplashScreen(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "success", gin.H{
		"latest_version": 2,
		"time":           time.Now().Format("15:04:05"),
	})
}

func VersionCheck(ctx *gin.Context) {
	app_version := ctx.DefaultQuery("app_version", "1")

	app_version_int, _ := strconv.Atoi(app_version)

	response.BaseResponse(ctx, http.StatusOK, true, "success", gin.H{
		"app_version": app_version_int,
		"api_version": 3,
	})
}
