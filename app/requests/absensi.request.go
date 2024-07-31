package requests

type AbsensiRequest struct {
	Status     string  `json:"status" binding:"required"`
	Lat        float32 `json:"lat" binding:"required"`
	Lng        float32 `json:"lng" binding:"required"`
	Image      string  `json:"image" binding:"required"`
	Keterangan string  `json:"keterangan"`
	Document   string  `json:"document"`
}
