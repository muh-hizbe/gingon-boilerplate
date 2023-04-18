package bootstrap

import (
	"gingon-boilerplate/configs"
	"gingon-boilerplate/configs/app_config"
	"gingon-boilerplate/configs/cors_config"
	"gingon-boilerplate/configs/logging_config"
	"gingon-boilerplate/database"
	"gingon-boilerplate/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	// Call init configuration
	configs.InitConfig()

	// Call connection database
	database.ConnectDatabase()

	// Logging in file default
	logging_config.DefaultLogging()

	// Run Gin/Other Engine Framework here with them routes
	app := gin.Default()
	app.Use(cors_config.CorsConfig())
	routes.InitRoute(app)
	app.Run(app_config.PORT)
}
