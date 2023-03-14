package routes

import (
	"gingon-boilerplate/controllers/home_controller"
	"gingon-boilerplate/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func v1Route(app *gin.RouterGroup) {
	route := app.Group("/v1")

	route.GET("/user", user_controller.Index)
	route.GET("/foo", home_controller.Index)
}
