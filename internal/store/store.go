package store

import (
	"context"

	"github.com/pejeio/blood-donate-locator-api/internal/types"
)

type Store interface {
	CreateLocation(ctx context.Context, loc types.CreateLocationRequest) (*types.Location, error)
	GetLocationById(ctx context.Context, id string) (types.Location, error)
	GetLocations(ctx context.Context, query types.FindLocationsRequest) ([]types.Location, error)
	DeleteLocation(ctx context.Context, id string) (int, error)
	CountLocations(ctx context.Context) (int64, error)
}
