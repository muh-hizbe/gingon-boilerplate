package user_controller

import (
	"gingon-boilerplate/database"
	// "gingon-boilerplate/models"
	"gingon-boilerplate/responses"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	user := new([]responses.UserResponse)

	// err := database.DB.Model(models.User{}).First(&user).Error
	err := database.DB.Table("users").Find(&user).Error
	if err != nil {
		log.Println(err.Error())
	}

	ctx.JSON(200, gin.H{
		"message": user,
	})
}

func GetById(ctx *gin.Context) {
	user := new(responses.UserResponse)
	userId := ctx.Param("id")

	// err := database.DB.Model(models.User{}).First(&user).Error
	err := database.DB.Table("users").Where("id = ?", userId).Find(&user).Error
	if err != nil {
		log.Println(err.Error())
	}

	if user.Email == "" {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": user,
	})
}
