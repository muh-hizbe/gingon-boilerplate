package logging_config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var defaultPath = "logs/gin.log"

func createLogFolderIfNotExist(path string) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Creating", dir, "directory...")

		err := os.MkdirAll(dir, 0644)
		if err != nil {
			fmt.Println("Fail to create ", dir, " directory", dir)
		} else {
			fmt.Println(dir, "directory created.")
		}
	}
}

func openOrCreateLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logFile, _ = os.Create(defaultPath)
	}
	return logFile, nil
}

func DefaultLogging(path ...string) {
	gin.DisableConsoleColor()

	if len(path) > 0 && path[0] != "" {
		defaultPath = path[0]
	}

	createLogFolderIfNotExist(defaultPath)
	f, _ := openOrCreateLogFile(defaultPath)
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
}
