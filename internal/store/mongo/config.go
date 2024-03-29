package mongo

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(ctx context.Context, c *configs.Config) (store.Store, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority&authSource=%s", c.DBUserName, c.DBUserPassword, c.DBHost, c.DBPort, c.DBName)
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := mClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Info().Msg("🚀 Connected successfully to the database")
	client := &Client{
		Database: mClient.Database(c.DBName),
	}
	return client, nil
}
