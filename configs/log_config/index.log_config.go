package log_config

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var defaultLogPath = "./logs/gin.log"

func createLogFolderIfNotExist(path string) {
	dir := filepath.Dir(path)

	log.Println("Creating directory")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0644)
		if err != nil {
			log.Println("Fail to create directory")
		} else {
			log.Println("Directory created")
		}
	}

}

func openOrCreateLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		var errorCreateFile error

		logFile, errorCreateFile := os.Create(path)

		if errorCreateFile != nil {
			log.Println("Canot create log file", errorCreateFile, logFile)
		}
	}

	return logFile, nil
}

func DefaultLogging(path ...string) {
	gin.DisableConsoleColor()

	if len(path) > 0 && path[0] != "" {
		defaultLogPath = path[0]
	}

	createLogFolderIfNotExist(defaultLogPath)
	f, _ := openOrCreateLogFile(defaultLogPath)

	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
}
