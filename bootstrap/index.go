package bootstrap

import (
	"gin-gorm/configs"
	"gin-gorm/configs/app_config"
	"gin-gorm/configs/cors_config"
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
	app.Use(cors_config.CorsConfig())
	database.ConnectDatabase()
	routes.InitRoute(app)
	app.Run(app_config.PORT)
}
