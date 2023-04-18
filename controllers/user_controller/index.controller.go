package user_controller

import (
	"fmt"
	"gingon-boilerplate/database"
	"gingon-boilerplate/models"
	"gingon-boilerplate/utils"
	"net/http"

	// "gingon-boilerplate/models"
	"gingon-boilerplate/requests"
	"gingon-boilerplate/responses"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetAll(ctx *gin.Context) {
	log.Println("this logging test")
	log.Println("ERR", "this logging test")
	log.Println("INFO", "this logging test")
	claimsData, isExist := ctx.Get("claimsData")
	log.Println("WARN", "this logging test")
	if isExist {
		claims := claimsData.(jwt.MapClaims)
		role := claims["role"]
		fmt.Println("role", role)
	}

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

func Create(ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	hashedPassword, errHash := utils.HashPassword(userReq.Password)
	if errHash != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := new(models.User)
	user.Name = userReq.Name
	user.Email = userReq.Email
	user.Address = userReq.Address
	user.Phone = userReq.Phone
	user.Password = hashedPassword

	err := database.DB.Table("users").Create(&user).Error
	if err != nil {
		log.Println(err.Error())
	}

	ctx.JSON(200, gin.H{
		"message": user,
	})
}
