package routes

import (
	"gingon-boilerplate/controllers/home_controller"
	"gingon-boilerplate/controllers/user_controller"
	"gingon-boilerplate/middleware"
	"gingon-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func v1Route(app *gin.RouterGroup) {
	route := app.Group("/v1")

	route.GET("/user", user_controller.GetAll)
	route.GET("/auth/user", middleware.AuthMiddleWare, user_controller.GetAll)
	route.GET("/user/:id", user_controller.GetById)
	route.POST("/user", user_controller.Create)
	route.GET("/foo", home_controller.Index)
	route.POST("/image", utils.HandleFileUpload, home_controller.StoreImage)
	route.POST("/image-handler", home_controller.StoreImageHandler)
	route.POST("/file", utils.HandleFileUpload, home_controller.StoreFile)
	route.GET("/auth/admin", middleware.AdminMiddleWare, user_controller.GetAll)
}
