package middleware

import (
	"encoding/json"
	"gin-gorm/app/model"
	"gin-gorm/app/response"
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
		response.BaseResponse(ctx, http.StatusUnauthorized, false, "Authorization header is missing or invalid", nil)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	redis_record, _ := database.RedisGet(ctx, "rcabs_"+token)

	var data_auth model.DataAuth
	err_unmarsal := json.Unmarshal([]byte(redis_record), &data_auth)
	if err_unmarsal == nil {
		ctx.Set("data_auth", data_auth)
		ctx.Next()
		return
	}

	data_karyawan := new(model.Karyawan)
	err_auth := database.DB.
		Table(TABLE_KARYAWAN).
		Where("karyawan_token = ?", token).
		First(&data_karyawan).Error

	if err_auth != nil {
		response.BaseResponse(ctx, http.StatusUnauthorized, false, "Authorization token invalid", nil)
		return
	}

	data_auth = model.DataAuth{
		KaryawanId:    data_karyawan.KaryawanId,
		KaryawanNama:  data_karyawan.KaryawanNama,
		KaryawanEmail: data_karyawan.KaryawanEmail,
		CabangId:      data_karyawan.CabangId,
		Role:          data_karyawan.Role,
	}

	database.RedisSet(ctx, "rcabs_"+token, data_auth, 3600)

	ctx.Set("data_auth", data_auth)
	ctx.Next()
}
