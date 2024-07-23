package requests

type EodRequest struct {
	Name string `json:"name" binding:"required"`
}
