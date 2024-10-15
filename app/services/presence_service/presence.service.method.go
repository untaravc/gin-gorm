package presence_service

import (
	"gin-gorm/app/models"
	"gin-gorm/app/requests"
	"gin-gorm/database"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const TABLE_ABSENSI = "absensi"

var STATUS_REGULAR = []string{
	"M", "A", "F", "MA", "MEET", "FT",
}

var STATUS_OUT = []string{
	"C", "O", "X", "S",
}

func (s *PresenceService) CheckDistance(cabang models.Cabang, absensi_request requests.AbsensiRequest) bool {
	if absensi_request.Status != "T" {
		distance := distance(cabang.CabangLatitude, cabang.CabangLongitude, absensi_request.Lat, absensi_request.Lng)
		return distance <= 300
	}
	return true
}

func (s *PresenceService) CheckStatus(dataReq requests.AbsensiRequest) bool {
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

func (s *PresenceService) AdditionalHour(cabang models.Cabang) int {
	switch cabang.CabangTimezone {
	case "WITA":
		return 1
	case "WIT":
		return 2
	default:
		return 0
	}
}

func (s *PresenceService) CheckTodayPresence(data_karyawan models.Karyawan) *models.Absensi {
	today := time.Now()
	formattedDate := today.Format("2006-01-02")

	layout := "2006-01-02"
	start, _ := time.Parse(layout, formattedDate)
	end := start.Add(24 * time.Hour)

	data_absensi := new(models.Absensi)
	err_absensi := database.DB.Table(TABLE_ABSENSI).
		Where("absensi_created_at >= ? AND absensi_created_at < ?", start, end).
		Where("karyawan_id = ?", data_karyawan.KaryawanId).
		First(&data_absensi).Error

	if err_absensi != nil && err_absensi != gorm.ErrRecordNotFound {
		return nil

	}
	return data_absensi
}

func (s *PresenceService) CheckManagement(data_cabang models.Cabang) bool {
	management_ids := []int{
		889, 15,
	}

	today := time.Now()

	// Get the weekday index (Sunday = 0, Monday = 1, ..., Saturday = 6)
	dayIndex := today.Weekday()

	if data_cabang.CabangType == "manajemen" {
		if dayIndex == 0 {
			if isIntInArray(management_ids, data_cabang.CabangId) {
				return false
			}
		}
	}

	return true
}

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

func isIntInArray(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func isStringInArray(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func (s *PresenceService) IsNeedApproval(dataReq requests.AbsensiRequest) bool {
	switch dataReq.Status {
	case "MEET":
	case "T":
	case "FT":
		return true
	}
	return false
}

func (s *PresenceService) IsPresenceRegular(status string) bool {
	return isStringInArray(STATUS_REGULAR, status)
}

func (s *PresenceService) IsPresenceOut(status string) bool {
	return isStringInArray(STATUS_OUT, status)
}

func (s *PresenceService) CheckLogout(ctx *gin.Context, data_absensi models.Absensi) {
	// absensi regular belum checkout
	if data_absensi.AbsensiCheckin.Valid && data_absensi.StatusAbsensi != "T" && data_absensi.AbsensiCheckout.Valid == false {
		ctx.JSON(400, gin.H{
			"status": false,
			"result": "Telah Presensi Regular dan Belum Checkout",
		})
		ctx.Abort()
	}

	// absensi 2 belum checkout

	// lembur belum checkout
	// lembur 2 belum checkout
	// sudah terisi semua
}

func (s *PresenceService) PresenceMap(dataAbsensi models.Absensi) models.AbsensiMaped {

	absensi_maped := models.AbsensiMaped{
		AbsensiStatus:    dataAbsensi.AbsensiStatus,
		AbsensiId:        dataAbsensi.AbsensiId,
		KaryawanId:       dataAbsensi.KaryawanId,
		CabangId:         dataAbsensi.CabangId,
		ApprovalId:       dataAbsensi.ApprovalId,
		JumlahJamKerja:   dataAbsensi.JumlahJamKerja,
		JumlahJamLembur:  dataAbsensi.JumlahJamLembur,
		Keterlambatan:    dataAbsensi.Keterlambatan,
		ApprovedDate:     dataAbsensi.ApprovedDate,
		ApprovedBy:       dataAbsensi.ApprovedBy,
		AbsensiBy:        dataAbsensi.AbsensiBy,
		Documement:       dataAbsensi.Document,
		AbsendiCreatedAt: dataAbsensi.AbsensiCreatedAt,
		AbsensiUpdatedAt: dataAbsensi.AbsensiUpdatedAt,
		// Absensi:          s.GeneralPresence(dataReq),
		// Absensi2:         s.GeneralPresence2(dataReq),
		// Lembur:           s.OvertimePresence(dataReq),
		// Lembur2:          s.OvertimePresence2(dataReq),
	}

	return absensi_maped
}
