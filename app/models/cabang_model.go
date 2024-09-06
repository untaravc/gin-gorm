package models

import "time"

type Cabang struct {
	CabangId        int       `json:"cabang_id" gorm:"primaryKey"`
	CabangNama      string    `json:"cabang_nama"`
	CabangNamaSpv   string    `json:"cabang_nama_spv"`
	CabangKota      string    `json:"cabang_kota"`
	CabangTimezona  string    `json:"cabang_timezone"`
	CabangLatitude  float64   `json:"cabang_latitude"`
	CabangLongitude float64   `json:"cabang_longitude"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}
