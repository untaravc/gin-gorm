package models

import (
	"database/sql"
	"time"
)

type Absensi struct {
	AbsensiId                int             `json:"absensi_id" gorm:"primaryKey"`
	KaryawanId               int             `json:"karyawan_id"`
	CabangId                 int             `json:"cabang_id"`
	ApprovalId               sql.NullInt32   `json:"approval_id"`
	AbsensiCheckin           sql.NullTime    `json:"absensi_checkin"`
	AbsensiCheckout          sql.NullTime    `json:"absensi_checkout"`
	AbsensiFotoCheckin       sql.NullString  `json:"absensi_foto_checkin"`
	AbsensiFotoCheckout      sql.NullString  `json:"absensi_foto_checkout"`
	AbsensiKeterangan        sql.NullString  `json:"absensi_keterangan"`
	AbsensiPulangKeterangan  sql.NullString  `json:"absensi_pulang_keterangan"`
	AbsensiCheckoutLat       sql.NullFloat64 `json:"absensi_checkout_lat"`
	AbsensiCheckoutLng       sql.NullFloat64 `json:"absensi_checkout_lng"`
	Absensi2Checkin          sql.NullTime    `json:"absensi2_checkin"`
	Absensi2Checkout         sql.NullTime    `json:"absensi2_checkout"`
	Absensi2FotoCheckin      sql.NullString  `json:"absensi2_foto_checkin"`
	Absensi2FotoCheckout     sql.NullString  `json:"absensi2_foto_checkout"`
	Absensi2Keterangan       sql.NullString  `json:"absensi2_keterangan"`
	Absensi2PulangKeterangan sql.NullString  `json:"absensi2_pulang_keterangan"`
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
	Latitude         sql.NullFloat64 `json:"latitude"`
	Longitude        sql.NullFloat64 `json:"longitude"`
	JumlahJamKerja   sql.NullInt32   `json:"jumlah_jam_kerja"`
	JumlahJamLembur  sql.NullInt32   `json:"jumlah_jam_lembur"`
	Keterlambatan    sql.NullInt32   `json:"keterlambatan"`
	AbsensiStatus    string          `json:"absensi_status"`
	AbsensiCreatedAt time.Time       `gorm:"column:absensi_created_at;autoCreateTime"`
	AbsensiUpdatedAt time.Time       `gorm:"column:absensi_updated_at;autoUpdateTime"`
	Approved         sql.NullBool    `json:"approved"`
	ApprovedBy       sql.NullInt32   `json:"approved_by"`
	ApprovedDate     sql.NullTime    `json:"approved_date"`
	AbsensiBy        sql.NullInt32   `json:"absensi_by"`
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
	StatusAbsensi2 sql.NullString `json:"status_absensi2"`
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
	Document  sql.NullString `json:"document"`
	Document2 sql.NullString `json:"document2"`
	// is_training_regular        int       `json:"is_training_regular"`
	// approved_note              int       `json:"approved_note"`
	ApiVersion sql.NullString `json:"api_version"`
}

type AbsensiMaped struct {
	AbsensiStatus    string           `json:"absensi_status"`
	AbsensiId        int              `json:"absensi_id"`
	KaryawanId       int              `json:"karyawan_id"`
	CabangId         int              `json:"cabang_id"`
	ApprovalId       *int32           `json:"approval_id"`
	JumlahJamKerja   *int32           `json:"jumlah_jam_kerja"`
	JumlahJamLembur  *int32           `json:"jumlah_jam_lembur"`
	Keterlambatan    *int32           `json:"keterlambatan"`
	ApprovedDate     *time.Time       `json:"approved_date"`
	ApprovedBy       *int32           `json:"approved_by"`
	AbsensiBy        *int32           `json:"absensi_by"`
	Documement       *string          `json:"document"`
	AbsendiCreatedAt time.Time        `json:"absensi_created_at"`
	AbsensiUpdatedAt time.Time        `json:"absensi_updated_at"`
	Absensi          GeneralPresence  `json:"absensi"`
	Absensi2         GeneralPresence  `json:"absensi2"`
	Lembur           OvertimePresence `json:"lembur"`
	Lembur2          OvertimePresence `json:"lembur2"`
}

type GeneralPresence struct {
	Status             string       `json:"status"`
	Checkin            sql.NullTime `json:"checkin"`
	CheckinFoto        string       `json:"checkin_foto"`
	CheckinKeterangan  string       `json:"checkin_keterangan"`
	CheckinLat         float64      `json:"checkin_lat"`
	CheckinLng         float64      `json:"checkin_lng"`
	Checkout           sql.NullTime `json:"checkout"`
	CheckoutFoto       string       `json:"checkout_foto"`
	CheckoutKeterangan string       `json:"checkout_keterangan"`
	CheckoutLat        float64      `json:"checkout_lat"`
	CheckoutLng        float64      `json:"checkout_lng"`
	Approved           string       `json:"approved"`
}

type OvertimePresence struct {
	Status             string       `json:"status"`
	Checkin            sql.NullTime `json:"checkin"`
	CheckinFoto        string       `json:"checkin_foto"`
	CheckinKeterangan  string       `json:"checkin_keterangan"`
	CheckinLat         float64      `json:"checkin_lat"`
	CheckinLng         float64      `json:"checkin_lng"`
	Checkout           sql.NullTime `json:"checkout"`
	CheckoutFoto       string       `json:"checkout_foto"`
	CheckoutKeterangan string       `json:"checkout_keterangan"`
	CheckoutLat        float64      `json:"checkout_lat"`
	CheckoutLng        float64      `json:"checkout_lng"`
	Approved           string       `json:"approved"`
	LemburStatus       string       `json:"lembur_status"`
	LemburAppovedAt    sql.NullTime `json:"lembur_approved_at"`
	LemburApprovedBy   int          `json:"lembur_approved_by"`
	Document           string       `json:"document"`
}
