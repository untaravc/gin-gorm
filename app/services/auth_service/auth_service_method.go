package auth_service

import (
	"gin-gorm/app/models"

	"github.com/gin-gonic/gin"
)

func (s *AuthService) GetAuth(ctx *gin.Context) models.Karyawan {
	data_auth, exists := ctx.Get("data_auth")

	if !exists {
		ctx.JSON(400, gin.H{"error": "Token Invalid"})
		ctx.Abort()
	}
	data_karyawan := data_auth.(*models.Karyawan)

	return *data_karyawan
}
