package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/sync/errgroup"
)

func (s *Server) CreateLocation(c *fiber.Ctx) error {
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

	newLocation, err := s.Store.CreateLocation(s.Ctx, *payload)

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

	limitOffset, err := GetPaginationQueryParams(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}
	findOptions := PaginationMongoOptions(limitOffset)

	g.Go(func() error {
		locs, err := s.Store.GetLocations(s.Ctx, bson.M{}, *findOptions)
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
		Data: locations,
		Meta: types.ResponseMeta{Count: locationsCount},
	})
}

func (s *Server) DeleteLocation(c *fiber.Ctx) error {
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

func (l *Server) UserIsLocationAdmin(c *fiber.Ctx) error {
	if can, _ := l.Enforcer.Enforce(c.Locals("_user"), "locations", "write"); !can {
		return c.Status(fiber.StatusForbidden).JSON(
			JsonErrorResponse{Message: "Forbidden"},
		)
	}
	return c.Next()
}
