package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coordinates struct {
	Latitude  float32 `json:"lat,omitempty" validate:"number"`
	Longitude float32 `json:"lng,omitempty" validate:"number"`
}

type Address struct {
	Street      string `json:"street,omitempty"`
	Number      int32  `json:"number,omitempty"`
	City        string `json:"city,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type Location struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Address     *Address           `bson:"address" json:"address"`
	Coordinates *Coordinates       `bson:"coordinates" json:"coordinates,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
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
	Name        string       `json:"name" validate:"required"`
	Coordinates *Coordinates `json:"coordinates" validate:"required"`
	Address     *Address     `json:"address"`
}

type PaginationRequest struct {
	Limit  string `query:"limit"`
	Offset string `query:"offset"`
}
