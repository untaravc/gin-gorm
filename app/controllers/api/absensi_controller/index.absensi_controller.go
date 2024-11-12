package absensi_controller

import (
	"database/sql"
	"gin-gorm/app/models"
	"gin-gorm/app/requests"
	"gin-gorm/app/services/auth_service"
	"gin-gorm/app/services/presence_service"
	"gin-gorm/database"
	"time"

	"github.com/gin-gonic/gin"
)

const TABLE = "absensi"
const TABLE_ABSENSI = "absensi"
const TABLE_KARYAWAN = "karyawans"
const TABLE_CABANG = "cabangs"

func Checkin(ctx *gin.Context) {
	presenceService := presence_service.NewPresenceService()
	authService := auth_service.NewAuthService()

	// Validasi Request
	dataReq := new(requests.AbsensiRequest)
	errReq := ctx.ShouldBind(dataReq)
	if errReq != nil {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": errReq.Error(),
		})
		return
	}

	// Cek status absensi
	valid_status := presenceService.CheckStatus(*dataReq)
	if !valid_status {
		ctx.JSON(400, gin.H{
			"status": false,
			"result": "ERR status",
		})
		return
	}

	// Get data karyawan login via auth_user
	data_karyawan := authService.GetAuth(ctx)

	// Cek status karyawan
	if data_karyawan.KaryawanStatus != "aktif" {
		ctx.JSON(400, gin.H{
			"status": false,
			"result": "Karyawan tidak aktif",
		})
		return
	}

	data_cabang := new(models.Cabang)
	err_cabang := database.DB.Table(TABLE_CABANG).
		Where("cabang_id = ?", data_karyawan.CabangId).
		First(&data_cabang).Error
	if err_cabang != nil {
		ctx.JSON(400, gin.H{
			"status": false,
			"result": "Karyawan tidak terdaftar di cabang",
		})
		return
	}

	// Penambahan Jam
	// additional_hour := presenceService.AdditionalHour(*data_cabang)

	// Cek Manajemen
	is_management := presenceService.CheckManagement(*data_cabang)

	if !is_management {
		ctx.JSON(400, gin.H{
			"status": false,
			"test":   "Tidak ada jadwal hari ini",
		})
	}

	// Cek jarak
	in_radius := presenceService.CheckDistance(*data_cabang, *dataReq)
	if !in_radius {
		ctx.JSON(400, gin.H{
			"status": false,
			"result": "Jarak terlalu jauh",
		})
		return
	}

	// Cek presensi hari ini
	data_absensi := presenceService.CheckTodayPresence(data_karyawan)

	is_need_approval := presenceService.IsNeedApproval(*dataReq)
	approved := sql.NullBool{Bool: false, Valid: false}
	if is_need_approval {
		approved = sql.NullBool{Bool: true, Valid: true}
	}

	// Jika belum presensi, maka akan membuat absensi
	if data_absensi == nil {
		store := new(models.Absensi)
		store.KaryawanId = data_karyawan.KaryawanId
		store.CabangId = data_karyawan.CabangId
		store.AbsensiCheckin = sql.NullTime{Time: time.Now(), Valid: true}
		store.Latitude = sql.NullFloat64{Float64: dataReq.Lat, Valid: true}
		store.Longitude = sql.NullFloat64{Float64: dataReq.Lng, Valid: true}
		store.AbsensiFotoCheckin = sql.NullString{String: dataReq.Image, Valid: true}
		store.StatusAbsensi = dataReq.Status
		store.AbsensiStatus = dataReq.Status
		store.AbsensiKeterangan = sql.NullString{String: dataReq.Keterangan}
		store.Document = sql.NullString{String: dataReq.Document}
		store.ApiVersion = sql.NullString{String: "gin.v1"}
		store.Approved = approved

		errDb := database.DB.Table(TABLE).Create(&store).Error
		if errDb != nil {
			ctx.JSON(400, gin.H{
				"status":  false,
				"message": "Fail to create data absesi",
				"error":   errDb,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"status": true,
			"result": "Checkin Success",
		})
		return
	}

	// Jika sudah presensi regular, tidak bisa masuk
	if presenceService.IsPresenceRegular(data_absensi.StatusAbsensi) {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": "Telah presensi masuk",
			"result":  presenceService.PresenceMap(*data_absensi),
		})
		return
	}

	// Jika sudah presensi out, tidak bisa masuk
	if presenceService.IsPresenceOut(data_absensi.AbsensiStatus) {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": "Anda telah presensi tidak masuk",
		})
		return
	}

	// Jika sudah presensi tidak masuk, tidak bisa masuk
	if presenceService.IsPresenceOut(dataReq.Status) {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": "Telah telah presensi masuk",
		})
		return
	}

	presenceService.CheckLogout(ctx, *data_absensi)

	// Belum absen regular
	if !data_absensi.AbsensiCheckin.Valid {
		data_absensi.StatusAbsensi = dataReq.Status
		data_absensi.Latitude = sql.NullFloat64{Float64: dataReq.Lat, Valid: true}
		data_absensi.Longitude = sql.NullFloat64{Float64: dataReq.Lng, Valid: true}
		data_absensi.AbsensiCheckin = sql.NullTime{Time: time.Now(), Valid: true}
		data_absensi.AbsensiFotoCheckin = sql.NullString{String: dataReq.Image, Valid: true}
		data_absensi.AbsensiKeterangan = sql.NullString{String: dataReq.Keterangan}
		data_absensi.Approved = approved

		database.DB.Table(TABLE).Where("absensi_id = ?", data_absensi.AbsensiId).Updates(data_absensi)

		ctx.JSON(200, gin.H{
			"status":  true,
			"message": "Berhasil melakukan presensi",
		})
		return
	}

	// jika sudah absen 1, sudah logout -> checkin2
	if data_absensi.AbsensiCheckout.Valid && !data_absensi.Absensi2Checkin.Valid {
		data_absensi.StatusAbsensi2 = sql.NullString{String: dataReq.Status, Valid: true}
		data_absensi.Absensi2Checkin = sql.NullTime{Time: time.Now(), Valid: true}
		data_absensi.AbsensiCheckoutLat = sql.NullFloat64{Float64: dataReq.Lat, Valid: true}
		data_absensi.AbsensiCheckoutLng = sql.NullFloat64{Float64: dataReq.Lat, Valid: true}
		data_absensi.Absensi2FotoCheckin = sql.NullString{String: dataReq.Image, Valid: true}
		data_absensi.Absensi2Keterangan = sql.NullString{String: dataReq.Keterangan}
		data_absensi.Approved = approved

		database.DB.Table(TABLE).Where("absensi_id = ?", data_absensi.AbsensiId).Updates(data_absensi)

		ctx.JSON(200, gin.H{
			"status":  true,
			"message": "Berhasil melakukan presensi",
		})
	}
	// jika sudah absen 2, sudah logout
	if data_absensi.Absensi2Checkout.Valid {
		ctx.JSON(400, gin.H{
			"status":  false,
			"message": "Telah presensi pulang",
			"result":  presenceService.PresenceMap(*data_absensi),
		})
	}

	ctx.JSON(400, gin.H{
		"status":  false,
		"message": "Telah presensi masuk",
		"result":  presenceService.PresenceMap(*data_absensi),
	})
}

// Clear Presence for Development Purpose
func ClearToday(ctx *gin.Context) {
	auth_service := auth_service.NewAuthService()

	data_karyawan := auth_service.GetAuth(ctx)

	today := time.Now()
	formattedDate := today.Format("2006-01-02")

	layout := "2006-01-02"
	start, _ := time.Parse(layout, formattedDate)
	end := start.Add(24 * time.Hour)

	data_absensi := new(models.Absensi)
	err_absensi := database.DB.Table(TABLE_ABSENSI).
		Where("absensi_created_at >= ? AND absensi_created_at < ?", start, end).
		Where("karyawan_id = ?", data_karyawan.KaryawanId).
		Delete(&data_absensi).Error

	if err_absensi != nil {
		ctx.JSON(500, gin.H{
			"status": false,
			"result": "ERR absensi",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  true,
		"message": "deleted",
		"result":  data_karyawan,
	})
}
