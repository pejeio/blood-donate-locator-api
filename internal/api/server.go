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
	log.Printf("👂 Listening and serving HTTP on %s\n", s.Config.ServerPort)
	log.Fatal(s.App.Listen(":" + s.Config.ServerPort))
}

func (s *Server) Routes() {
	// Locations
	router := s.App.Group("locations")
	router.Get("/", s.FindLocations)
	router.Post("/", BasicAuthHandler(), s.UserIsLocationAdmin, s.CreateLocation)
	router.Delete("/:id", BasicAuthHandler(), s.UserIsLocationAdmin, s.DeleteLocation)
}

func (s *Server) Cors() {
	s.App.Use(CorsHandler())
}
