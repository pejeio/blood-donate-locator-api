package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/auth"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/store"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Ctx        context.Context
	App        *fiber.App
	Config     *configs.Config
	AuthClient *auth.Client
	Store      store.Store
}

func NewServer(ctx context.Context, c *configs.Config, s store.Store, authC *auth.Client, app *fiber.App) *Server {
	return &Server{
		Ctx:        ctx,
		App:        app,
		Config:     c,
		AuthClient: authC,
		Store:      s,
	}
}

func (s *Server) Start() {
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

	// Middleware that checks if the token is valid and sets the user ID to the Fiber Locals
	authMiddleware := Protect(s.AuthClient)

	// Define routes
	locationRouter.Get("/", authMiddleware, s.FindLocations)
	locationRouter.Get("/lookup", s.FindLocationsByCoordinates)
	locationRouter.Get("/:id", s.FindLocation)
	locationRouter.Post("/", authMiddleware, CreateLocationAuthzMiddleware(s.AuthClient), s.CreateLocation)
	locationRouter.Delete("/:id", DeleteLocationAuthzMiddleware(s.AuthClient), s.DeleteLocation)
}

func (s *Server) Cors() {
	s.App.Use(CorsHandler())
}
