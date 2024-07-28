package routes

import (
	"gin-gorm/app/controllers/base_controller"
	"gin-gorm/app/controllers/eods_controller"
	"gin-gorm/app/controllers/file_controller"
	"gin-gorm/app/middleware"
	"gin-gorm/configs/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app

	route.GET("/", base_controller.Index)
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	route.POST("/upload", file_controller.HandleUploadFile)

	eodRoute := route.Group("", middleware.AuthMiddleware)
	eodRoute.GET("/eods", eods_controller.GetAllData)
	eodRoute.GET("/eods/:id", eods_controller.GetById)
	eodRoute.POST("/eods", eods_controller.Store)
	eodRoute.PATCH("/eods/:id", eods_controller.Update)
	eodRoute.DELETE("/eods/:id", eods_controller.Delete)
}
