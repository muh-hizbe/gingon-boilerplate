package cors_config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var origins = []string{"*"}

func CorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = origins	
	return cors.New(config)
}
