package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
)

func CreateLocation(c *fiber.Ctx) error {
	payload := new(types.CreateLocationRequest)
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
	newLocation := types.Location{
		Name:        payload.Name,
		Coordinates: payload.Coordinates,
		Address:     payload.Address,
	}
	result := configs.Db().Create(&newLocation)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: result.Error.Error()},
		)
		return result.Error
	}
	return c.Status(fiber.StatusCreated).JSON(newLocation)
}

func FindLocations(c *fiber.Ctx) error {
	pag := GetPaginationQueryParams(c)
	var locations []types.Location
	var totalCount int64

	configs.Db().Find(&locations).Count(&totalCount)
	configs.Db().
		Limit(pag.Limit).
		Offset(pag.Offset).
		Order("created_at DESC").
		Find(&locations)

	return c.JSON(types.ResponseWithPagination{
		Data: locations,
		Meta: types.ResponseMeta{Count: totalCount},
	})
}
