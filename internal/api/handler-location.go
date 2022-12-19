package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
)

func (s *Server) LocationCollection() *mongo.Collection {
	return s.client.Database(s.config.DBName).Collection("locations")
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
	doc, err := s.LocationCollection().InsertOne(c.Context(), newLocation)

	newLocation.ID = doc.InsertedID.(primitive.ObjectID)

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

	var totalCount int64
	locations := make([]types.Location, 0)

	g.Go(func() error {
		count, err := s.LocationCollection().CountDocuments(c.Context(), filter)
		totalCount = count
		return err
	})

	g.Go(func() error {
		cursor, err := s.LocationCollection().Find(c.Context(), filter, opts)
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

func (s *Server) DeleteLocation(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			JsonErrorResponse{Message: "Invalid ID"},
		)
	}
	filter := bson.M{"_id": objId}
	res, _ := s.LocationCollection().DeleteOne(context.Background(), filter)
	if res.DeletedCount > 0 {
		return c.SendStatus(200)
	}
	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			JsonErrorResponse{Message: "Location not found"},
		)
	}

	return c.SendStatus(fiber.StatusInternalServerError)
}

func (l *Server) UserIsLocationAdmin(c *fiber.Ctx) error {
	if can, _ := l.enforcer.Enforce(c.Locals("_user"), "locations", "write"); !can {
		return c.Status(fiber.StatusForbidden).JSON(
			JsonErrorResponse{Message: "Forbidden"},
		)
	}
	return c.Next()
}
