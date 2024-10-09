package routes

import (
	"gin-gorm/app/controllers/api/absensi_controller"
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

	// eodRoute := route.Group("", middleware.AuthMiddleware)
	// eodRoute.GET("/eods", eods_controller.GetAllData)
	// eodRoute.GET("/eods/:id", eods_controller.GetById)
	// eodRoute.POST("/eods", eods_controller.Store)
	// eodRoute.PATCH("/eods/:id", eods_controller.Update)
	// eodRoute.DELETE("/eods/:id", eods_controller.Delete)

	api := route.Group("/api", middleware.AuthMiddleware)
	// Absensi
	api.POST("/check-in", absensi_controller.Checkin)

	api.GET("/report-cabang", middleware.AuthMiddleware, report_controller.Index)
}
