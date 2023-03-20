package app_config

import (
	"log"
	"os"
)

var PORT = ":8000" //string

// for public asset access
var PUBLIC_ROUTE = "/public"
var PUBLIC_ASSETS_DIR = "./public"

func InitAppConfig() {
	env_APP_PORT := os.Getenv("APP_PORT")
	if env_APP_PORT != "" {
		log.Println("APP_PORT => ", env_APP_PORT)
		PORT = env_APP_PORT
	}
}