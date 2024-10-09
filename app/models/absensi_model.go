package models

import (
	"database/sql"
	"time"
)

type Absensi struct {
	AbsensiId               int             `json:"absensi_id" gorm:"primaryKey"`
	KaryawanId              int             `json:"karyawan_id"`
	CabangId                int             `json:"cabang_id"`
	ApprovalId              int             `json:"approval_id"`
	AbsensiCheckin          sql.NullTime    `json:"absensi_checkin"`
	AbsensiCheckout         sql.NullTime    `json:"absensi_checkout"`
	AbsensiFotoCheckin      string          `json:"absensi_foto_checkin"`
	AbsensiFotoCheckout     sql.NullString  `json:"absensi_foto_checkout"`
	AbsensiKeterangan       string          `json:"absensi_keterangan"`
	AbsensiPulangKeterangan sql.NullString  `json:"absensi_pulang_keterangan"`
	AbsensiCheckoutLat      sql.NullFloat64 `json:"absensi_checkout_lat"`
	AbsensiCheckoutLng      sql.NullFloat64 `json:"absensi_checkout_lng"`
	// Absensi2Checkin          sql.NullTime `json:"absensi2_checkin"`
	// Absensi2Checkout         sql.NullTime `json:"absensi2_checkout"`
	// Absensi2FotoCheckin      string       `json:"absensi2_foto_checkin"`
	// Absensi2FotoCheckout     string       `json:"absensi2_foto_checkout"`
	// Absensi2Keterangan       string       `json:"absensi2_keterangan"`
	// Absensi2PulangKeterangan string       `json:"absensi2_pulang_keterangan"`
	// lembur_checkin             int       `json:"lembur_checkin"`
	// lembur_checkin_lat         int       `json:"lembur_checkin_lat"`
	// lembur_checkin_lng         int       `json:"lembur_checkin_lng"`
	// lembur_checkout            int       `json:"lembur_checkout"`
	// lembur_checkout_lat        int       `json:"lembur_checkout_lat"`
	// lembur_checkout_lng        int       `json:"lembur_checkout_lng"`
	// lembur_foto_checkin        int       `json:"lembur_foto_checkin"`
	// lembur_foto_checkout       int       `json:"lembur_foto_checkout"`
	// lembur_keterangan          int       `json:"lembur_keterangan"`
	// lembur_pulang_keterangan   int       `json:"lembur_pulang_keterangan"`
	Latitude  sql.NullFloat64 `json:"latitude"`
	Longitude sql.NullFloat64 `json:"longitude"`
	// jumlah_jam_kerja           int       `json:"jumlah_jam_kerja"`
	// jumlah_jam_lembur          int       `json:"jumlah_jam_lembur"`
	// keterlambatan              int       `json:"keterlambatan"`
	AbsensiStatus    string        `json:"absensi_status"`
	AbsensiCreatedAt time.Time     `gorm:"column:absensi_created_at;autoCreateTime"`
	AbsensiUpdatedAt time.Time     `gorm:"column:absensi_updated_at;autoUpdateTime"`
	Approved         sql.NullBool  `json:"approved"`
	ApprovedBy       sql.NullInt32 `json:"approved_by"`
	ApprovedDate     sql.NullTime  `json:"approved_date"`
	AbsensiBy        sql.NullInt32 `json:"absensi_by"`
	// lembur_status              int       `json:"lembur_status"`
	// lembur2_status             int       `json:"lembur2_status"`
	// lembur_approved_at         int       `json:"lembur_approved_at"`
	// lembur_approved_by         int       `json:"lembur_approved_by"`
	// lembur2_approved_at        int       `json:"lembur2_approved_at"`
	// lembur2_approved_by        int       `json:"lembur2_approved_by"`
	// absensi_approved           int       `json:"absensi_approved"`
	// absensi2_approved          int       `json:"absensi2_approved"`
	StatusAbsensi string `json:"status_absensi"`
	// absensi2_checkin_lat       int       `json:"absensi2_checkin_lat"`
	// absensi2_checkin_lng       int       `json:"absensi2_checkin_lng"`
	// absensi2_checkout_lat      int       `json:"absensi2_checkout_lat"`
	// absensi2_checkout_lng      int       `json:"absensi2_checkout_lng"`
	// status_absensi2            int       `json:"status_absensi2"`
	// Lembur2Checkin          sql.NullTime `json:"lembur2_checkin"`
	// Lembur2FotoCheckin      string       `json:"lembur2_foto_checkin"`
	// Lembur2Keterangan       string       `json:"lembur2_keterangan"`
	// Lembur2CheckinLat       float64      `json:"lembur2_checkin_lat"`
	// Lembur2CheckinLng       float64      `json:"lembur2_checkin_lng"`
	// Lembur2Checkout         sql.NullTime `json:"lembur2_checkout"`
	// Lembur2FotoCheckout     string       `json:"lembur2_foto_checkout"`
	// Lembur2CheckoutLat      float64      `json:"lembur2_checkout_lat"`
	// Lembur2CheckoutLng      float64      `json:"lembur2_checkout_lng"`
	// Lembur2PulangKeterangan string       `json:"lembur2_pulang_keterangan"`
	Document sql.NullString `json:"document"`
	// Document2               string       `json:"document2"`
	// is_training_regular        int       `json:"is_training_regular"`
	// approved_note              int       `json:"approved_note"`
	ApiVersion sql.NullString `json:"api_version"`
}
