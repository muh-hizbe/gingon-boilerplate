package middleware

import (
	"gingon-boilerplate/constanta"
	"gingon-boilerplate/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func AdminMiddleWare(ctx *gin.Context) {
	token := ctx.GetHeader("x-token")
	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println("err :: ", err)
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	role := claims["role"].(string)
	if role != constanta.ADMIN {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "forbidden access",
		})
		return
	}

	ctx.Next()
}
