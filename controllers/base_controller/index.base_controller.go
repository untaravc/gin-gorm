package base_controller

import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
		"message": "go web app",
	})
}
