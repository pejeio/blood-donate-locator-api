package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func BasicAuthHandler() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "123456",
		},
		ContextUsername: "_user",
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(
				JsonErrorResponse{Message: "Unauthorized"},
			)
		},
	})
}

func UserIsLocationWriter(c *fiber.Ctx) error {
	if can, _ := Enforcer.Enforce(c.Locals("_user"), "locations", "write"); !can {
		return c.Status(fiber.StatusForbidden).JSON(
			JsonErrorResponse{Message: "Forbidden"},
		)
	}
	return c.Next()
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
