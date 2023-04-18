package logging_config

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var defaultPath = "logs/gin.log"

func openOrCreateLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logFile, _ = os.Create("path")
	}
	return logFile, nil
}

func DefaultLogging(path ...string) {
	gin.DisableConsoleColor()

	if len(path) > 0 && path[0] != "" {
		defaultPath = path[0]
	}

	f, _ := openOrCreateLogFile(defaultPath)
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
}
