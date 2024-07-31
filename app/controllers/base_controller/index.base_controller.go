package base_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
		"message": "go web app",
	})
}

func View(ctx *gin.Context) {
	data := gin.H{
		"title":  "Halaman",
		"number": 2,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
