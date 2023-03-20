package routes

import (
	"gingon-boilerplate/configs/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	// Static asset or public asset
	app.Static(app_config.PUBLIC_ROUTE, app_config.PUBLIC_ASSETS_DIR)

	// routes
	api := app.Group("/api")
	v1Route(api)
}
