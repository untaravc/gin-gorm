package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token != "Bearer Token" {
		ctx.JSON(401, gin.H{
			"success": false,
			"message": "Unauthenticateds",
		})
		ctx.Abort()
	}

	ctx.Set("name", "Untara")
	ctx.Next()
}
