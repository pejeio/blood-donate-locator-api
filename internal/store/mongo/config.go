package mongo

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"github.com/pejeio/blood-donate-locator-api/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(c *configs.Config, ctx context.Context) (store.Store, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", c.DBUserName, c.DBUserPassword, c.DBHost, c.DBPort)
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := mClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("🚀 Connected successfully to the database")
	client := &Client{
		Database: mClient.Database(c.DBName),
	}
	return client, nil
}
