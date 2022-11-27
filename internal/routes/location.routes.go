package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/controllers"
	"github.com/pejeio/blood-donate-locator-api/internal/middlewares"
)

type LocationRouteController struct {
	locationController controllers.LocationController
}

func NewRouteLocationController(locationController controllers.LocationController) LocationRouteController {
	return LocationRouteController{locationController}
}

func (lc *LocationRouteController) LocationRoute(app *fiber.App) {
	router := app.Group("locations")
	router.
		Get("/", lc.locationController.FindLocations)
	router.
		Use(middlewares.BasicAuthHandler()).
		Post("/", lc.locationController.CreateLocation)
}
