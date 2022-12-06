package api

import (
	"github.com/gofiber/fiber/v2"
)

func LocationRoutes(app *fiber.App) {
	router := app.Group("locations")

	router.Get("/", FindLocations)

	router.
		Use(BasicAuthHandler()).
		Use(UserIsLocationWriter).
		Post("/", CreateLocation)
}
