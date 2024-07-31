package absensi_controller

import (
	"gin-gorm/app/models"
	"gin-gorm/app/requests"
	"gin-gorm/app/utils"
	"gin-gorm/database"

	"github.com/gin-gonic/gin"
)

const TABLE = "absensi"

func Index(ctx *gin.Context) {

}

func Create(ctx *gin.Context) {
	dataReq := new(requests.AbsensiRequest)

	errReq := ctx.ShouldBind(dataReq)

	if errReq != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": errReq.Error(),
			"error":   utils.ParseErrorMessages(errReq.Error()),
		})
		return
	}

	store := new(models.Absensi)
	store.KaryawanId = 2

	errDb := database.DB.Table(TABLE).Create(&store).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "fail to create data",
			"error":   errDb,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"success": true,
	})
}

func Update(ctx *gin.Context) {

}
