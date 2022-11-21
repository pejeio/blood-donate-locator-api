package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/models"
	"gorm.io/gorm"
)

type LocationController struct {
	DB *gorm.DB
}

func NewLocationController(DB *gorm.DB) LocationController {
	return LocationController{DB}
}

func (lc *LocationController) CreateLocation(c *fiber.Ctx) error {
	payload := new(models.CreateLocationRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
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
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
		return result.Error
	}
	c.Status(fiber.StatusCreated).JSON(newLocation)
	return nil
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

	c.JSON(models.ResponseWithPagination{
		Data: locations,
		Meta: models.ResponseMeta{Count: totalCount},
	})
	return nil
}
