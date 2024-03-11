package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func (s *Server) CreateLocation(c *fiber.Ctx) error {
	log.Info().Msg("Creating location")

	// Parse the request body
	body := new(types.CreateLocationRequest)
	if err := c.BodyParser(body); err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			JSONErrorResponse{Message: err.Error()},
		)
	}

	// Validate the request body
	if errors := ValidateStruct(*body); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Set the created_by field
	body.CreatedBy = GetUserIDFromCtx(c)

	// Create the location
	newLocation, err := s.Store.CreateLocation(s.Ctx, *body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JSONErrorResponse{Message: err.Error()},
		)
	}

	// Return the created location
	return c.Status(fiber.StatusCreated).JSON(newLocation)
}

// FindLocations retrieves a list of locations based on the provided query parameters.
func (s *Server) FindLocations(c *fiber.Ctx) error {
	pagQParams, err := GetPaginationQueryParams(c)
	if err != nil {
		return err
	}

	query := types.FindLocationsRequest{
		City:       c.Query("city"),
		PostalCode: c.Query("postal_code"),
		Limit:      pagQParams.Limit,
		Offset:     pagQParams.Offset,
	}

	locations, totalCount, err := s.Store.GetLocations(s.Ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JSONErrorResponse{Message: err.Error()},
		)
	}

	return c.JSON(types.ResponseWithPagination{
		Response: types.Response{
			Data: locations,
		},
		Meta: types.ResponseMeta{Count: totalCount},
	})
}

// FindLocationsByCoordinates finds locations by coordinates.
func (s *Server) FindLocationsByCoordinates(c *fiber.Ctx) error {
	var (
		g         errgroup.Group
		locations []types.Location
	)

	lat, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONErrorResponse{Message: "could not parse latitude"},
		)
	}

	lng, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONErrorResponse{Message: "could not parse longitude"},
		)
	}

	maxDist, err := strconv.Atoi(c.Query("max_distance"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JSONErrorResponse{Message: "could not parse max_distance"},
		)
	}

	query := types.LookupLocationRequest{
		Coordinates: &types.Coordinates{
			Latitude:  lat,
			Longitude: lng,
		},
		MaxDistance: maxDist,
	}

	g.Go(func() error {
		locs, err := s.Store.ReverseGeoCodeLocations(s.Ctx, query)
		locations = locs
		return err
	})
	if err := g.Wait(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JSONErrorResponse{Message: err.Error()},
		)
	}

	return c.JSON(types.Response{
		Data: locations,
	})
}

// FindLocation finds the location based on the provided ID.
func (s *Server) FindLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	loc, err := s.Store.GetLocationByID(s.Ctx, id)
	if err != nil {
		if err.Error() == "not found" {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			JSONErrorResponse{Message: err.Error()},
		)
	}
	return c.JSON(loc)
}

// DeleteLocation deletes a location.
func (s *Server) DeleteLocation(c *fiber.Ctx) error {
	log.Info().Msg("Deleting location")
	id := c.Params("id")
	delCount, err := s.Store.DeleteLocation(s.Ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JSONErrorResponse{Message: err.Error()},
		)
	}
	if delCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			JSONErrorResponse{Message: "Location not found"},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"deleted": delCount},
	)
}
