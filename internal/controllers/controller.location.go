package controllers

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LocationController struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func NewLocationController(DB *gorm.DB, Enforcer *casbin.Enforcer) LocationController {
	return LocationController{DB, Enforcer}
}

func (lc *LocationController) CreateLocation(c *fiber.Ctx) error {
	if can, _ := lc.Enforcer.Enforce(c.Locals("_user"), "locations", "write"); !can {
		return c.Status(fiber.StatusUnauthorized).JSON(
			JsonErrorResponse{Message: "Unauthorized"},
		)
	}
	payload := new(models.CreateLocationRequest)

	if err := c.BodyParser(payload); err != nil {
		log.Errorln(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}

	errors := ValidateStruct(*payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newLocation := models.Location{
		Name:        payload.Name,
		Coordinates: payload.Coordinates,
		Address:     payload.Address,
	}
	result := lc.DB.Create(&newLocation)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: result.Error.Error()},
		)
		return result.Error
	}
	return c.Status(fiber.StatusCreated).JSON(newLocation)
}

func (lc *LocationController) FindLocations(c *fiber.Ctx) error {
	pag := GetPaginationQueryParams(c)
	var locations []models.Location
	var totalCount int64

	lc.DB.Find(&locations).Count(&totalCount)
	lc.DB.
		Limit(pag.Limit).
		Offset(pag.Offset).
		Order("created_at DESC").
		Find(&locations)

	return c.JSON(models.ResponseWithPagination{
		Data: locations,
		Meta: models.ResponseMeta{Count: totalCount},
	})
}
