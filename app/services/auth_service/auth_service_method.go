package auth_service

import (
	"fmt"
	"gin-gorm/app/model"
	"gin-gorm/app/response"

	"github.com/gin-gonic/gin"
)

func (s *AuthService) GetAuth(ctx *gin.Context) model.DataAuth {
	data_auth, exists := ctx.Get("data_auth")
	fmt.Println("get ctx", data_auth)
	if !exists {
		response.BaseResponse(ctx, 400, true, "auth_service.GetAuth: Token Invalid", nil)
		ctx.Abort()
	}

	result := data_auth.(model.DataAuth)

	return result
}
