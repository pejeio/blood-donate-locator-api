package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/controllers"
	"github.com/pejeio/blood-donate-locator-api/internal/middlewares"
	"github.com/pejeio/blood-donate-locator-api/internal/models"
	"github.com/pejeio/blood-donate-locator-api/internal/routes"
	log "github.com/sirupsen/logrus"
)

var (
	app *fiber.App

	LocationController      controllers.LocationController
	LocationRouteController routes.LocationRouteController
)

func main() {
	log.Info("🛫 Starting the app")

	// Config
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("🧐 Could not load environment variables", err)
	}

	// DB
	configs.ConnectDB(&config)
	models.AutoMigrate()

	// Server
	app = fiber.New(fiber.Config{DisableStartupMessage: true})
	middlewares.UseCors(app)

	// Controllers
	LocationController = controllers.NewLocationController(configs.Db())
	LocationRouteController = routes.NewRouteLocationController(LocationController)

	// Routes
	LocationRouteController.LocationRoute(app)

	log.Printf("👂 Listening and serving HTTP on %s\n", config.ServerPort)
	log.Fatal(app.Listen(":" + config.ServerPort))
}
