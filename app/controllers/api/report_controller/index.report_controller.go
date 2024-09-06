package report_controller

import (
	"fmt"
	"gin-gorm/app/models"
	"gin-gorm/database"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const TABLE_ABSENSI = "absensi"
const TABLE_KARYAWAN = "karyawan"
const TABLE_CABANG = "cabangs"

type QueryFilter struct {
	// defining values of struct
	from      string
	to        string
	cabang_id int
}

type ReportMonth struct {
	date string
	// absensi models.Absensi
}

var data_base_absensi []models.Absensi
var data_base_karyawan *[]models.Karyawan
var data_base_cabang *models.Cabang

func getAbsensi(wg *sync.WaitGroup, query QueryFilter, ctx *gin.Context) {

	data_absensi := new([]models.Absensi)
	err_absensi := database.DB.Table(TABLE_ABSENSI).Where("absensi_created_at > ?", query.from).
		Where("cabang_id = ?", query.cabang_id).
		Find(&data_absensi).Error

	if err_absensi != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"result":  "ERR absensi",
		})
		return
	}
	wg.Done()
	data_base_absensi = *data_absensi
}

func getKaryawan(wg *sync.WaitGroup, ctx *gin.Context, data_absensi []models.Absensi) {
	data_karyawan := new([]models.Karyawan)

	unique_karyawan := getUniqueKaryawan(data_absensi)
	karyawan_ids := make([]int, 0, len(unique_karyawan))

	for _, kry := range unique_karyawan {
		karyawan_ids = append(karyawan_ids, kry.KaryawanId)
	}

	err_karyawan := database.DB.Table(TABLE_KARYAWAN).Where("karyawan_id IN ?", karyawan_ids).
		Find(&data_karyawan).Error

	if err_karyawan != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"result":  "ERR karyawan",
		})
		return
	}

	data_base_karyawan = data_karyawan
	wg.Done()
}

func getCabang(wg *sync.WaitGroup, query QueryFilter, ctx *gin.Context) {
	data_cabang := new(models.Cabang)
	err_absensi := database.DB.Table(TABLE_CABANG).
		Where("cabang_id = ?", query.cabang_id).
		First(&data_cabang).Error

	if err_absensi != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"result":  "ERR absensi",
		})
		return
	}
	wg.Done()
	data_base_cabang = data_cabang
}

func endOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func Index(ctx *gin.Context) {
	report_date := []map[string]string{}

	query := QueryFilter{}
	if ctx.Query("from") != "" {
		query.from = ctx.Query("from")
	} else {
		query.from = time.Now().Format("YYYY-MM-DD")
	}

	if ctx.Query("to") != "" {
		query.to = ctx.Query("to")
	} else {
		query.to = time.Now().Format("YYYY-MM-DD")
	}

	if ctx.Query("cabang_id") != "" {
		num, err := strconv.Atoi(ctx.Query("cabang_id"))
		if err == nil {
			query.cabang_id = num
		}
	}

	from, err := time.Parse("2006-01-02", query.from)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	last_day := endOfMonth(from)
	date_count := last_day.Day()

	for i := 0; i < date_count; i++ {
		prefix := "0"

		if i > 8 {
			prefix = ""
		}

		newEntry := map[string]string{"date": query.from[0:7] + "-" + prefix + strconv.Itoa(i+1)}
		report_date = append(report_date, newEntry)
	}
	fmt.Print(report_date)

	ctx.JSON(200, gin.H{
		"last":   last_day,
		"result": report_date,
	})

	return

	var wg sync.WaitGroup

	wg.Add(1)
	go getAbsensi(&wg, query, ctx)
	wg.Wait()

	wg.Add(1)
	go getKaryawan(&wg, ctx, data_base_absensi)
	wg.Wait()

	wg.Add(1)
	go getCabang(&wg, query, ctx)
	wg.Wait()

	ctx.JSON(200, gin.H{
		"success": true,
		"query":   query.from,
		"result":  data_base_karyawan,
		"cabang":  data_base_cabang,
	})
}

func getUniqueKaryawan(absensi []models.Absensi) []models.Absensi {
	uniqueMap := make(map[int]models.Absensi)

	for _, abs := range absensi {
		uniqueMap[abs.KaryawanId] = abs
	}

	uniqueList := make([]models.Absensi, 0, len(uniqueMap))
	for _, abs := range uniqueMap {
		uniqueList = append(uniqueList, abs)
	}

	return uniqueList
}
