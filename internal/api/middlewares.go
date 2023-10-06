package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
)

func BasicAuthHandler() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users:           configs.GetAuthUsers(),
		ContextUsername: "_user",
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(
				JSONErrorResponse{Message: "Unauthorized"},
			)
		},
	})
}

func CorsHandler() fiber.Handler {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}
