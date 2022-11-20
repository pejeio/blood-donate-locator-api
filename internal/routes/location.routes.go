package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pejeio/blood-donate-locator-api/internal/controllers"
)

type LocationRouteController struct {
	locationController controllers.LocationController
}

func NewRouteLocationController(locationController controllers.LocationController) LocationRouteController {
	return LocationRouteController{locationController}
}

func (lc *LocationRouteController) LocationRoute(rg *gin.RouterGroup) {
	router := rg.Group("locations")
	router.GET("/", lc.locationController.FindLocations)
	router.POST("/", lc.locationController.CreateLocation)
}
