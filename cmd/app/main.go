package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/api"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/store/mongo"
	log "github.com/sirupsen/logrus"
)

func main() {
	configs.SetUpLogging()
	log.Info("🛫 Starting the app")

	// Config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("🧐 Could not load environment variables", err)
	}

	// Init context
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	defer cancelFunc()

	// Database
	dbClient, err := mongo.Init(ctx, &cfg)
	if err != nil {
		log.Error("❌ Failed to connect to the database", err)
	}

	// Authorization
	enforcer, err := api.NewEnforcer(&cfg)
	if err != nil {
		log.Error("❌ Failed to set up authorization", err)
	}

	// Fiber App
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// Server
	server := api.NewServer(ctx, &cfg, dbClient, enforcer, app)
	server.Start()
}
