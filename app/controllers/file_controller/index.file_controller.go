package file_controller

import (
	"fmt"
	"gin-gorm/app/utils"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const TABLE = "eods"

func HandleUploadFile(ctx *gin.Context) {
	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"success": true,
			"message": "file required",
		})
		return
	}

	extendsionFile := filepath.Ext(fileHeader.Filename)
	time := time.Now().UTC().Format("20240130")
	fileName := fmt.Sprintf("%s%s", utils.RandomString(10), extendsionFile)

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s/%s", time, fileName))

	if errUpload != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}
