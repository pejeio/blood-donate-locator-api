package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/api"
	"github.com/pejeio/blood-donate-locator-api/internal/auth"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/store/mongo"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set up logging
	configs.SetUpLogging()

	log.Info("🛫 Starting the app")

	// Load config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("🧐 Could not load environment variables", err)
	}

	// Set up context with signal handling
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	defer cancelFunc()

	// Initialize database client
	dbClient, err := mongo.Init(ctx, &cfg)
	if err != nil {
		log.Error("❌ Failed to connect to the database", err)
	}

	// Initialize authentication client
	authClient := auth.NewClient(cfg.KCBaseURL, cfg.KCClientID, cfg.KCClientSecret, cfg.KCRealm)
	if err != nil {
		log.Error("❌ Failed to set up authentication client", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Create API server
	server := api.NewServer(ctx, &cfg, dbClient, authClient, app)

	// Start server
	server.Start()
}
