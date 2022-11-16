package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/models"
	log "github.com/sirupsen/logrus"
)

var (
	server *gin.Engine
)

func main() {
	log.Info("🛫 Starting the app")

	// Config
	configs.SetUpLogging()
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// DB
	configs.ConnectDB(&config)
	models.AutoMigrate()

	// Server
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()
	log.Fatal(server.Run(":" + config.ServerPort))
}
