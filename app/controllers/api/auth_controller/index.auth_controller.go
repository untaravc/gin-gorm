package auth_controller

import (
	"gin-gorm/app/model"
	"gin-gorm/app/request"
	"gin-gorm/app/response"
	"gin-gorm/app/services/auth_service"
	"gin-gorm/app/utils"
	"gin-gorm/database"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var body request.LoginRequest

	// Validasi Request
	if err := ctx.ShouldBind(&body); err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Email atau Password Tidak Sesuai")
		return
	}

	// Get Karyawan
	karyawan, err := getKaryawan(body)
	if err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Karyawan tidak terdaftar")
		return
	}

	// Validasi Password
	if !checkPassword(karyawan.KaryawanPassword, body.Password) {
		hashed, err := hashPassword(body.Password)
		if err != nil {
			response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Password Salah: "+body.Password)
			return
		}
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Password Salah: "+hashed)
		return
	}

	// Create Token, Save to DB
	new_token := utils.RandomString(64)

	if err := updateToken(karyawan, new_token); err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Gagal")
		return
	}

	// Save to Redis expired 1 Hour
	karyawan_redis := model.DataAuth{
		KaryawanId:    karyawan.KaryawanId,
		KaryawanNama:  karyawan.KaryawanNama,
		KaryawanEmail: karyawan.KaryawanEmail,
		CabangId:      karyawan.CabangId,
		Role:          karyawan.Role,
	}

	err_redis := database.RedisSet(ctx, "rcabs_"+new_token, karyawan_redis, 3600)

	if err_redis != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", err_redis)
		return
	}

	karyawan.KaryawanPassword = nil
	karyawan.KaryawanToken = &new_token
	response.BaseResponse(ctx, 200, true, "success", karyawan)
}

func Auth(ctx *gin.Context) {
	authService := auth_service.NewAuthService()
	data_auth := authService.GetAuth(ctx)

	response.BaseResponse(ctx, http.StatusOK, true, "OK", data_auth)
}

func UpdatePassword(ctx *gin.Context) {
	var body request.UpdatePasswordRequest

	// Validasi Request
	if err := ctx.ShouldBind(&body); err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Email atau Password Tidak Sesuai")
		return
	}

	// Get Data Auth from token
	authService := auth_service.NewAuthService()
	data_auth := authService.GetAuth(ctx)

	// Get Karyawan by email
	login_request := request.LoginRequest{
		Email:    *data_auth.KaryawanEmail,
		Password: body.OldPassword,
	}

	karyawan, err := getKaryawan(login_request)
	if err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Karyawan tidak terdaftar")
		return
	}

	// Validasi Password
	if !checkPassword(karyawan.KaryawanPassword, body.OldPassword) {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Password Lama Salah")
		return
	}

	// Update Password
	hashed, err := hashPassword(body.NewPassword)
	if err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "GAGAL")
		return
	}

	err = database.DB.Table("karyawan").
		Where("karyawan_email = ?", *data_auth.KaryawanEmail).
		Update("karyawan_password", hashed).Error

	if err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "error", "Gagal memperbarui password")
		return
	} else {
		response.BaseResponse(ctx, http.StatusOK, true, "success", nil)
	}
}

func checkPassword(hashedPassword *string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte((*hashedPassword)), []byte(password))
	return err == nil
}

func getKaryawan(login_request request.LoginRequest) (model.Karyawan, error) {
	data_karyawan := new(model.Karyawan)
	err_karyawan := database.DB.Table("karyawan").
		Where("karyawan_email = ?", login_request.Email).
		First(&data_karyawan).Error

	if err_karyawan != nil {
		return model.Karyawan{}, err_karyawan
	} else {
		return *data_karyawan, nil
	}
}

func updateToken(data_karyawan model.Karyawan, token string) error {
	err_karyawan := database.DB.Table("karyawan").
		Where("karyawan_email = ?", data_karyawan.KaryawanEmail).
		Update("karyawan_token", token).Error

	if err_karyawan != nil {
		return err_karyawan
	} else {
		return nil
	}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
