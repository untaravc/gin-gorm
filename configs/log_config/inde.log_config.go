package log_config

import (
	"os"
	"path/filepath"
)

var defaultLogPath = "./logs/gin.log"

func createLogFolderIfNotExist(path string) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0644)
	}

}
