package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	app      *fiber.App
	config   *configs.Config
	client   *mongo.Client
	enforcer *casbin.Enforcer
}

func NewServer(config *configs.Config, client *mongo.Client, enf *casbin.Enforcer, app *fiber.App) *Server {
	return &Server{
		app:      app,
		config:   config,
		client:   client,
		enforcer: enf,
	}
}

func (s *Server) Start() {
	configs.InitAuthUsers()
	s.Cors()
	s.Routes()
	log.Printf("👂 Listening and serving HTTP on %s\n", s.config.ServerPort)
	log.Fatal(s.app.Listen(":" + s.config.ServerPort))
}

func (s *Server) Routes() {
	// Locations
	router := s.app.Group("locations")
	router.Get("/", s.FindLocations)
	router.Post("/", BasicAuthHandler(), s.UserIsLocationAdmin, s.CreateLocation)
	router.Delete("/:id", BasicAuthHandler(), s.UserIsLocationAdmin, s.DeleteLocation)
}

func (s *Server) Cors() {
	s.app.Use(CorsHandler())
}
