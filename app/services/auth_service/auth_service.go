package auth_service

import "github.com/gin-gonic/gin"

type AuthService struct {
	ctx *gin.Context
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
