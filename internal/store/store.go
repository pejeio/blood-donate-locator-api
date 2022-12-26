package store

import (
	"context"

	"github.com/pejeio/blood-donate-locator-api/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	CreateLocation(ctx context.Context, loc types.CreateLocationRequest) (*types.Location, error)
	GetLocations(ctx context.Context, filter bson.M, findOptions options.FindOptions) ([]types.Location, error)
	DeleteLocation(ctx context.Context, id string) (int, error)
	CountLocations(ctx context.Context) (int64, error)
}
