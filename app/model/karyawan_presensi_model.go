package model

type KaryawanPresensi struct {
	KaryawanId   int
	KaryawanNama string
	PresensiList []DatePresensi
}

type DatePresensi struct {
	Date            string
	KaryawanAbsensi Absensi
}
