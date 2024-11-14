package model

import "time"

type Cabang struct {
	CabangId        int       `json:"cabang_id" gorm:"primaryKey"`
	CabangNama      string    `json:"cabang_nama"`
	CabangNamaSpv   string    `json:"cabang_nama_spv"`
	CabangKota      string    `json:"cabang_kota"`
	CabangTimezone  string    `json:"cabang_timezone"`
	CabangLatitude  float64   `json:"cabang_latitude"`
	CabangLongitude float64   `json:"cabang_longitude"`
	CabangType      string    `json:"cabang_type"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}
