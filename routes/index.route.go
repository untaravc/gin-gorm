package routes

import (
	"gin-gorm/controllers/book_controller"
	"gin-gorm/controllers/eods_controller"
	"gin-gorm/controllers/karyawan_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app

	route.GET("/karyawan", karyawan_controller.GetAllUser)
	route.GET("/karyawan/:id", karyawan_controller.GetById)
	route.GET("/book", book_controller.GetAllBook)

	route.GET("/eods", eods_controller.GetAllData)
	route.GET("/eods/:id", eods_controller.GetById)
	route.POST("/eods", eods_controller.Store)
	route.PATCH("/eods/:id", eods_controller.Update)
	route.DELETE("/eods/:id", eods_controller.Delete)
}
