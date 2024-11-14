package request

type AbsensiRequest struct {
	Status     string  `json:"status" binding:"required"`
	Lat        float64 `json:"lat" binding:"required"`
	Lng        float64 `json:"lng" binding:"required"`
	Image      string  `json:"image" binding:"required"`
	Keterangan string  `json:"keterangan"`
	Document   string  `json:"document"`
	KaryawanId string  `json:"karyawan_id"`
	CabangId   string  `json:"cabang_id"`
}
