package mongo

import (
	"context"
	"errors"

	"github.com/pejeio/blood-donate-locator-api/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Database *mongo.Database
}

func (c *Client) LocationsCollection() *mongo.Collection {
	return c.Database.Collection("locations")
}

func (c *Client) GetLocations(ctx context.Context, filter bson.M, findOptions options.FindOptions) ([]types.Location, error) {
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	locations := make([]types.Location, 0)

	cursor, err := c.LocationsCollection().Find(ctx, filter, &findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var location types.Location
		err := cursor.Decode(&location)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func (c *Client) CreateLocation(ctx context.Context, loc types.CreateLocationRequest) (*types.Location, error) {
	newLocation := types.Location{
		Name:        loc.Name,
		Coordinates: loc.Coordinates,
		Address:     loc.Address,
	}
	doc, err := c.LocationsCollection().InsertOne(ctx, newLocation)
	if err != nil {
		return nil, err
	}
	newLocation.ID = doc.InsertedID.(primitive.ObjectID)

	_, err = newLocation.MarshalBSON()
	if err != nil {
		return nil, err
	}

	return &newLocation, err
}

func (c *Client) DeleteLocation(ctx context.Context, id string) (int, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, errors.New("Invalid ID")
	}
	filter := bson.M{"_id": objId}
	res, err := c.LocationsCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(res.DeletedCount), nil
}

func (c *Client) CountLocations(ctx context.Context) (int64, error) {
	filter := bson.D{}

	return c.LocationsCollection().CountDocuments(ctx, filter)
}
