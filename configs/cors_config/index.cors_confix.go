package cors_config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"}

	return cors.New(config)
}
