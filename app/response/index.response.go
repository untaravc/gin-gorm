package response

import (
	"github.com/gin-gonic/gin"
)

type TypeResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func BaseResponse(c *gin.Context, httpStatus int, success bool, message string, result interface{}) {
	c.JSON(httpStatus, TypeResponse{
		Success: success,
		Message: message,
		Result:  result,
	})
}
