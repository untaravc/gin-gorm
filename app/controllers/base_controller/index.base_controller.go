package base_controller

import (
	"gin-gorm/app/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response.BaseResponse(ctx, http.StatusOK, true, "OK", "Hello World")
}

func View(ctx *gin.Context) {
	data := gin.H{
		"title":  "Halaman",
		"number": 2,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
