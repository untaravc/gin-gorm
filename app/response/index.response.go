package response

import (
	"github.com/gin-gonic/gin"
)

type TypeResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func BaseResponse(ctx *gin.Context, httpStatus int, success bool, message string, result interface{}) {
	ctx.AbortWithStatusJSON(httpStatus, TypeResponse{
		Success: success,
		Message: message,
		Result:  result,
	})
}
