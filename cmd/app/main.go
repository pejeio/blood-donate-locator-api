package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/api"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	log "github.com/sirupsen/logrus"
)

var (
	app *fiber.App
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
	configs.AutoMigrate()

	// Authz
	api.NewEnforcer(config)

	// Server
	app = fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(api.CorsHandler())

	// Routes
	api.LocationRoutes(app)

	log.Printf("👂 Listening and serving HTTP on %s\n", config.ServerPort)
	log.Fatal(app.Listen(":" + config.ServerPort))
}
