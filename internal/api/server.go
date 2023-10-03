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

func NewServer(c *configs.Config, s store.Store, enf *casbin.Enforcer, app *fiber.App, ctx context.Context) *Server {
	return &Server{
		App:      app,
		Config:   c,
		Store:    s,
		Enforcer: enf,
		Ctx:      ctx,
	}
}

func (s *Server) Start() {
	err := configs.InitAuthUsers()
	if err != nil {
		log.Println(err)
	}
	s.Cors()
	s.Routes()
	s.Store.CreateLocationIndexes(s.Ctx)
	log.Printf("👂 Listening and serving HTTP on %s\n", s.Config.ServerPort)
	log.Fatal(s.App.Listen(":" + s.Config.ServerPort))
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
