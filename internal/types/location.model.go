package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coordinates struct {
	Latitude  float64 `json:"lat,omitempty" validate:"number"`
	Longitude float64 `json:"lng,omitempty" validate:"number"`
}

type GeoJSONPoint struct {
	Type        string     `bson:"type" json:"type"`
	Coordinates [2]float64 `bson:"coordinates" json:"coordinates"`
}

type Address struct {
	Street      string `json:"street,omitempty"`
	Number      int32  `json:"number,omitempty"`
	City        string `json:"city,omitempty"`
	PostalCode  string `json:"postal_code,omitempty" bson:"postal_code"`
	CountryCode string `json:"country_code,omitempty" bson:"country_code"`
}

type Location struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Address   *Address           `bson:"address" json:"address"`
	Geometry  *GeoJSONPoint      `bson:"geometry" json:"geometry,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (l *Location) MarshalBSON() ([]byte, error) {
	if l.CreatedAt.IsZero() {
		l.CreatedAt = time.Now()
	}
	l.UpdatedAt = time.Now()

	type my Location
	return bson.Marshal((*my)(l))
}

type CreateLocationRequest struct {
	Name        string       `json:"name" validate:"required,min=5"`
	Coordinates *Coordinates `json:"coordinates" validate:"required"`
	Address     *Address     `json:"address"`
}

type FindLocationsRequest struct {
	PostalCode string
	City       string
	Limit      int
	Offset     int
}

type LookupLocationRequest struct {
	*Coordinates
	MaxDistance int `json:"max_distance"`
}

type PaginationRequest struct {
	Limit  string `query:"limit"`
	Offset string `query:"offset"`
}
