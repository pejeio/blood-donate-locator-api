package api

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/store"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	App      *fiber.App
	Config   *configs.Config
	Store    store.Store
	Enforcer *casbin.Enforcer
	Ctx      context.Context
}

func NewServer(ctx context.Context, c *configs.Config, s store.Store, enf *casbin.Enforcer, app *fiber.App) *Server {
	return &Server{
		Ctx:      ctx,
		App:      app,
		Config:   c,
		Store:    s,
		Enforcer: enf,
	}
}

func (s *Server) Start() {
	// Initialize authentication users
	if err := configs.InitAuthUsers(); err != nil {
		log.Println(err)
	}

	// Set up CORS and routes
	s.Cors()
	s.Routes()

	// Create location indexes
	if err := s.Store.CreateLocationIndexes(s.Ctx); err != nil {
		log.Fatal(err)
	}

	// Start the server
	serverAddr := ":" + s.Config.ServerPort
	log.Printf("👂 Listening and serving HTTP on %s\n", serverAddr)
	log.Fatal(s.App.Listen(serverAddr))
}

func (s *Server) Routes() {
	// Group routes for "locations"
	locationRouter := s.App.Group("locations")

	// Middleware for routes requiring Basic Authentication
	authMiddleware := BasicAuthHandler()

	// Define routes
	locationRouter.Get("/", s.FindLocations)
	locationRouter.Get("/lookup", s.FindLocationsByCoordinates)
	locationRouter.Get("/:id", s.FindLocation)
	locationRouter.Post("/", authMiddleware, s.UserIsLocationAdmin, s.CreateLocation)
	locationRouter.Delete("/:id", authMiddleware, s.UserIsLocationAdmin, s.DeleteLocation)
}

func (s *Server) Cors() {
	s.App.Use(CorsHandler())
}
