package bootstrap

import (
	"gin-gorm/configs"
	"gin-gorm/configs/app_config"
	"gin-gorm/configs/cors_config"
	"gin-gorm/configs/log_config"
	"gin-gorm/database"
	"gin-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	configs.InitConfig()

	app := gin.Default()

	// Add cors config
	app.Use(cors_config.CorsConfig())

	database.ConnectDatabase()
	database.InitRedisClient()

	log_config.DefaultLogging()
	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
