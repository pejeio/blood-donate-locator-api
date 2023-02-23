package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func (s *Server) CreateLocation(c *fiber.Ctx) error {
	log.Println("Creating location")
	body := new(types.CreateLocationRequest)
	if err := c.BodyParser(body); err != nil {
		log.Errorln(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}
	errors := ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newLocation, err := s.Store.CreateLocation(s.Ctx, *body)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}
	return c.Status(fiber.StatusCreated).JSON(newLocation)
}

func (s *Server) FindLocations(c *fiber.Ctx) error {
	var (
		g              errgroup.Group
		locations      []types.Location
		locationsCount int64
	)

	pagQParams, _ := GetPaginationQueryParams(c)

	query := types.FindLocationsRequest{
		City:       c.Query("city"),
		PostalCode: c.Query("postal_code"),
		Limit:      pagQParams.Limit,
		Offset:     pagQParams.Offset,
	}

	g.Go(func() error {
		locs, err := s.Store.GetLocations(s.Ctx, query)
		locations = locs
		return err
	})

	g.Go(func() error {
		count, err := s.Store.CountLocations(s.Ctx)
		locationsCount = count
		return err
	})

	if err := g.Wait(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}

	return c.JSON(types.ResponseWithPagination{
		Response: types.Response{
			Data: locations,
		},
		Meta: types.ResponseMeta{Count: locationsCount},
	})
}

func (s *Server) FindLocationsByCoordinates(c *fiber.Ctx) error {
	var (
		g         errgroup.Group
		locations []types.Location
	)

	lat, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JsonErrorResponse{Message: "could not parse latitude"},
		)
	}

	lng, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JsonErrorResponse{Message: "could not parse longitude"},
		)
	}

	maxDist, err := strconv.Atoi(c.Query("max_distance"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JsonErrorResponse{Message: "could not parse max_distance"},
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
			JsonErrorResponse{Message: err.Error()},
		)
	}

	return c.JSON(types.Response{
		Data: locations,
	})
}

func (s *Server) FindLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	loc, err := s.Store.GetLocationById(s.Ctx, id)
	if err != nil {
		if err.Error() == "not found" {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}
	return c.JSON(loc)
}

func (s *Server) DeleteLocation(c *fiber.Ctx) error {
	log.Println("Deleting location")
	id := c.Params("id")
	delCount, err := s.Store.DeleteLocation(s.Ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}
	if delCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			JsonErrorResponse{Message: "Location not found"},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{"deleted": delCount},
	)
}

func (s *Server) UserIsLocationAdmin(c *fiber.Ctx) error {
	if can, _ := s.Enforcer.Enforce(c.Locals("_user"), "locations", "write"); !can {
		return c.Status(fiber.StatusForbidden).JSON(
			JsonErrorResponse{Message: "Forbidden"},
		)
	}
	return c.Next()
}
