package routes

import (
	"gin-gorm/controllers/book_controller"
	"gin-gorm/controllers/karyawan_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app

	route.GET("/karyawan", karyawan_controller.GetAllUser)
	route.GET("/book", book_controller.GetAllBook)
}
