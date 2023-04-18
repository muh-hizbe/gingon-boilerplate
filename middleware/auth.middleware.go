package middleware

import (
	"gingon-boilerplate/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "invalid token",
		})
		return
	}
	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	// token := ctx.GetHeader("x-token")
	// if token == "" {
	// 	ctx.AbortWithStatusJSON(401, gin.H{
	// 		"message": "unauthenticated",
	// 	})
	// 	return
	// }

	claimsData, err := utils.DecodeToken(token)
	if err != nil {
		log.Println("err :: ", err)
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated",
		})
		return
	}

	ctx.Set("claimsData", claimsData)

	ctx.Next()
}
