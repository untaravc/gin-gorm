package middleware

import (
	"gin-gorm/app/models"
	"gin-gorm/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const TABLE_KARYAWAN = "karyawan"

func AuthMiddleware(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")

	// Check if the Authorization header is empty or not a Bearer token
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or invalid"})
		ctx.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	data_auth := new(models.Karyawan)
	err_auth := database.DB.
		Table(TABLE_KARYAWAN).
		Where("karyawan_token = ?", token).
		First(&data_auth).Error

	if err_auth != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalid"})
		ctx.Abort()
		return
	}

	ctx.Set("data_auth", data_auth)
	ctx.Next()
}
