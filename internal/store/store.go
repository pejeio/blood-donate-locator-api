package store

import (
	"context"

	"github.com/pejeio/blood-donate-locator-api/internal/types"
)

type Store interface {
	CreateLocationIndexes(ctx context.Context) error
	CreateLocation(ctx context.Context, loc types.CreateLocationRequest) (*types.Location, error)
	GetLocationByID(ctx context.Context, id string) (types.Location, error)
	GetLocations(ctx context.Context, query types.FindLocationsRequest) ([]types.Location, int64, error)
	DeleteLocation(ctx context.Context, id string) (int, error)
	ReverseGeoCodeLocations(ctx context.Context, query types.LookupLocationRequest) ([]types.Location, error)
}
