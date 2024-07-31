package models

import "time"

type Absensi struct {
	AbsensiId        int       `json:"absensi_id" gorm:"primaryKey"`
	KaryawanId       int       `json:"karyawan_id"`
	CabangId         int       `json:"cabang_id"`
	AbsensiCreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}
