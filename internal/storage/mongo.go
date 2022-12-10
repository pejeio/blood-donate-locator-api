package storage

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pejeio/blood-donate-locator-api/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(config *configs.Config) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("🚀 Connected successfully to the database")
	return client, nil
}

func DisconnectDb(client *mongo.Client) error {
	log.Println("🚀 Disconnected the database")
	return client.Disconnect(context.TODO())
}
