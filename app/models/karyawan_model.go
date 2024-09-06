package models

import "time"

type Karyawan struct {
	KaryawanId   int        `json:"karyawan_id" gorm:"primaryKey"`
	KaryawanNama *string    `json:"karyawan_nama"`
	CreatedAt    *time.Time `json:"created_at"`
}
