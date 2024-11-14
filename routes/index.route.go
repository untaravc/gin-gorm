package routes

import (
	"gin-gorm/app/controllers/api/auth_controller"
	"gin-gorm/app/controllers/api/home_controller"
	"gin-gorm/app/controllers/api/report_controller"
	"gin-gorm/app/controllers/base_controller"
	"gin-gorm/app/middleware"
	"gin-gorm/configs/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
	route.Static("/css", "../templates/css")
	route.LoadHTMLGlob("templates/pages/*.html")

	route.GET("/", base_controller.Index)
	// route.GET("/view", base_controller.View)

	// route.POST("/upload", file_controller.HandleUploadFile)
	// route.GET("/test", test_controller.ConnectToRedis)

	api := route.Group("/api")
	api.POST("/login", auth_controller.Login)
	api.GET("/auth", middleware.AuthMiddleware, auth_controller.Auth)

	// Absensi
	// api.POST("/check-in", absensi_controller.Checkin)
	// api.DELETE("/presence", absensi_controller.ClearToday)

	api.GET("/report-cabang", middleware.AuthMiddleware, report_controller.Index)

	// api.POST("/splash-screen", home_controller.SplashScreen) // ok
	api.GET("/version-check", home_controller.VersionCheck)

}
