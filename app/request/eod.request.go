package request

type EodRequest struct {
	Name string `json:"name" binding:"required"`
}
