package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

func (s *Server) LocationsCollection() *mongo.Collection {
	return s.mongoClient.Database(s.config.DBName).Collection("locations")
}

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
	newLocation := types.Location{
		Name:        payload.Name,
		Coordinates: payload.Coordinates,
		Address:     payload.Address,
	}
	newLocation.MarshalBSON()
	_, err := s.LocationsCollection().InsertOne(c.Context(), newLocation)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			JsonErrorResponse{Message: err.Error()},
		)
	}
	return c.Status(fiber.StatusCreated).JSON(newLocation)
}

func (s *Server) FindLocations(c *fiber.Ctx) error {
	pag := GetPaginationQueryParams(c)
	filter := bson.D{}
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetLimit(int64(pag.Limit))
	opts.SetSkip(int64(pag.Offset))

	g := new(errgroup.Group)

	var (
		totalCount int64
		locations  []types.Location
	)

	g.Go(func() error {
		count, err := s.LocationsCollection().CountDocuments(c.Context(), filter)
		totalCount = count
		return err
	})

	g.Go(func() error {
		cursor, err := s.LocationsCollection().Find(c.Context(), filter, opts)
		for cursor.Next(c.Context()) {
			var location types.Location
			err := cursor.Decode(&location)
			if err != nil {
				return err
			}
			locations = append(locations, location)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(types.ResponseWithPagination{
		Data: locations,
		Meta: types.ResponseMeta{Count: totalCount},
	})
}

func (l *Server) UserIsLocationWriter(c *fiber.Ctx) error {
	if can, _ := l.enforcer.Enforce(c.Locals("_user"), "locations", "write"); !can {
		return c.Status(fiber.StatusForbidden).JSON(
			JsonErrorResponse{Message: "Forbidden"},
		)
	}
	return c.Next()
}
