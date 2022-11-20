package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/controllers"
	"github.com/pejeio/blood-donate-locator-api/internal/models"
	"github.com/pejeio/blood-donate-locator-api/internal/routes"
	log "github.com/sirupsen/logrus"
)

var (
	server *gin.Engine

	LocationController      controllers.LocationController
	LocationRouteController routes.LocationRouteController
)

func main() {
	log.Info("🛫 Starting the app")

	// Config
	configs.SetUpLogging()
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("🧐 Could not load environment variables", err)
	}

	// DB
	configs.ConnectDB(&config)
	models.AutoMigrate()

	LocationController = controllers.NewLocationController(configs.Db())
	LocationRouteController = routes.NewRouteLocationController(LocationController)

	// Server
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()
	server.Use(configs.CorsHandleFunc())
	router := server.Group("/api")

	LocationRouteController.LocationRoute(router)

	log.Printf("👂 Listening and serving HTTP on %s\n", config.ServerPort)
	log.Fatal(server.Run(":" + config.ServerPort))
}
