package middleware

import (
	"gingon-boilerplate/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(ctx *gin.Context) {
	token := ctx.GetHeader("x-token")
	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	_, err := utils.DecodeToken(token)
	if err != nil {
		log.Println("err :: ", err)
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	ctx.Next()
}
