package configs

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandleFunc() gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"http://localhost"}
	cfg.AllowCredentials = true
	return cors.New(cfg)
}
