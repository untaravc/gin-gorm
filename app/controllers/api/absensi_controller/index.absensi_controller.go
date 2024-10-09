package absensi_controller

import (
	"database/sql"
	"gin-gorm/app/models"
	"gin-gorm/app/requests"
	"gin-gorm/app/utils"
	"gin-gorm/database"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const TABLE = "absensi"
const TABLE_ABSENSI = "absensi"
const TABLE_KARYAWAN = "karyawans"
const TABLE_CABANG = "cabangs"

func Checkin(ctx *gin.Context) {
	// Validasi Request
	dataReq := new(requests.AbsensiRequest)
	errReq := ctx.ShouldBind(dataReq)

	if errReq != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": errReq.Error(),
			"error":   utils.ParseErrorMessages(errReq.Error()),
		})
		return
	}

	data_auth, exists := ctx.Get("data_auth")
	if !exists {
		ctx.JSON(400, gin.H{"error": "Token Invalid"})
		return
	}

	// Cek status absensi
	valid_status := checkStatus(*dataReq)

	if !valid_status {
		ctx.JSON(400, gin.H{
			"success": false,
			"result":  "ERR status",
		})
		return
	}

	data_karyawan := data_auth.(*models.Karyawan)

	today := time.Now()
	formattedDate := today.Format("2006-01-02")

	layout := "2006-01-02"
	start, _ := time.Parse(layout, formattedDate)
	end := start.Add(24 * time.Hour)

	// Cek apakah sudah memiliki absensi
	data_absensi := new(models.Absensi)
	err_absensi := database.DB.Table(TABLE_ABSENSI).
		Where("absensi_created_at >= ? AND absensi_created_at < ?", start, end).
		Where("karyawan_id = ?", data_karyawan.KaryawanId).
		First(&data_absensi).Error

	data_cabang := new(models.Cabang)
	err_cabang := database.DB.Table(TABLE_CABANG).
		Where("cabang_id = ?", data_karyawan.CabangId).
		First(&data_cabang).Error

	if err_cabang != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"result":  "ERR cabang",
		})
		return
	}

	in_radius := checkDistance(*data_cabang, *dataReq)
	if !in_radius {
		ctx.JSON(400, gin.H{
			"success": false,
			"result":  "Jarak terlalu jauh",
		})
		return
	}

	// Cek apakah sudah melakukan absensi
	if err_absensi != nil && err_absensi == gorm.ErrRecordNotFound {
		store := new(models.Absensi)
		store.KaryawanId = data_karyawan.KaryawanId
		store.CabangId = data_karyawan.CabangId
		store.AbsensiCheckin = sql.NullTime{Time: time.Now(), Valid: true}
		store.Latitude = sql.NullFloat64{Float64: dataReq.Lat, Valid: true}
		store.Longitude = sql.NullFloat64{Float64: dataReq.Lng, Valid: true}
		store.AbsensiFotoCheckin = dataReq.Image
		store.StatusAbsensi = dataReq.Status
		store.AbsensiStatus = dataReq.Status
		store.AbsensiKeterangan = dataReq.Keterangan
		store.Document = sql.NullString{String: dataReq.Document}
		store.ApiVersion = sql.NullString{String: "gin.v1"}

		errDb := database.DB.Table(TABLE).Create(&store).Error
		if errDb != nil {
			ctx.JSON(400, gin.H{
				"success": false,
				"message": "Fail to create data absesi",
				"error":   errDb,
			})
			return
		}

		// Create new Absensi
		ctx.JSON(200, gin.H{
			"success": true,
			"result":  "Checkin Success",
		})
		return
	}

	ctx.JSON(400, gin.H{
		"success": true,
		"message": "Telah presensi masuk",
		"result":  "Checkin Success",
	})
}

func checkDistance(cabang models.Cabang, absensi_request requests.AbsensiRequest) bool {
	distance := distance(cabang.CabangLatitude, cabang.CabangLongitude, absensi_request.Lat, absensi_request.Lng)
	return distance <= 300
}

// create function to check the distance between origin lat and lng and destination lat and lng
func distance(lat1, lon1, lat2, lon2 float64) float64 {

	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	// haversine formula
	r = 6371000 // metres
	x := math.Sin((lo2-lo1)/2)*math.Sin((lo2-lo1)/2) + math.Cos(lo1)*math.Cos(lo2)*math.Sin((la2-la1)/2)*math.Sin((la2-la1)/2)
	c := 2 * math.Atan2(math.Sqrt(x), math.Sqrt(1-x))
	d := r * c // in metres
	return d
}

func checkStatus(dataReq requests.AbsensiRequest) bool {
	statusList := []string{
		"M",
		"MA",
		"F",
		"A",
		"MEET",
		"TC",
		"FT",
		"C",
		"O",
		"X",
		"alpha",
	}

	for _, status := range statusList {
		if status == dataReq.Status {
			return true
		}
	}
	return false
}
