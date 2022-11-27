package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/authz"
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
	Enforcer                *casbin.Enforcer
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
	app.Use(middlewares.CorsHandler())

	// Authorization
	Enforcer = authz.NewEnforcer("casbin.conf", "casbin_policy.csv")

	// Controllers
	LocationController = controllers.NewLocationController(configs.Db(), Enforcer)
	LocationRouteController = routes.NewRouteLocationController(LocationController)

	// Routes
	LocationRouteController.LocationRoute(app)

	log.Printf("👂 Listening and serving HTTP on %s\n", config.ServerPort)
	log.Fatal(app.Listen(":" + config.ServerPort))
}
