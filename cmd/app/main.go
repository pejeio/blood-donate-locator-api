package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/api"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/storage"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("🛫 Starting the app")

	// Config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("🧐 Could not load environment variables", err)
	}

	// Database
	mongoClient, err := storage.ConnectDB(&cfg)
	if err != nil {
		log.Fatal("❌ Failed to connect to the database", err)
	}
	defer func() {
		if err = storage.DisconnectDb(mongoClient); err != nil {
			panic(err)
		}
	}()

	// Authorization
	enforcer, err := api.NewEnforcer(&cfg)
	if err != nil {
		log.Fatal("❌ Failed to set up authorization", err)
	}

	// Fiber App
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// Server
	server := api.NewServer(&cfg, mongoClient, enforcer, app)
	server.Start()
}
