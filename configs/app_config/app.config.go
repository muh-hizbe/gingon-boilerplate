package app_config

import (
	"log"
	"os"
)

var PORT = ":8000" //string

func InitAppConfig() {
	env_APP_PORT := os.Getenv("APP_PORT")
	if env_APP_PORT != "" {
		log.Println("APP_PORT => ", env_APP_PORT)
		PORT = env_APP_PORT
	}
}