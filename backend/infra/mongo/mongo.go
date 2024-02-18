package infra

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database = "auto"
const collection = "replacements"

type MongoDriver struct {
	routerHost string
	logger     *slog.Logger

	client *mongo.Client
}

func NewMongoDriver(logger *slog.Logger, routerHost string) *MongoDriver {
	return &MongoDriver{
		routerHost: routerHost,
		logger:     logger,
	}
}

func (m *MongoDriver) getClient() (*mongo.Client, error) {
	if m.client != nil {
		return m.client, nil
	}

	url := "mongodb://" + m.routerHost + ":27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		m.logger.Error("could not establish a connection to %s.\n", slog.String("url", url))
		return nil, err
	}

	m.client = client
	return client, nil
}

func (m *MongoDriver) Close() error {
	if m.client != nil {
		if err := m.client.Disconnect(context.Background()); err != nil {
			m.logger.Error("could not properly disconnect from mongodb.\n")
			return err
		}
	}
	return nil
}

func (m *MongoDriver) Process(unit func(*mongo.Collection, *context.Context) error) error {

	client, err := m.getClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err != nil {
		return err
	}

	databaseHandler := client.Database(database)
	collectionHandler := databaseHandler.Collection(collection)

	err = unit(collectionHandler, &ctx)

	if err != nil {
		m.logger.Warn("could not process the mongodb operation.")
	}

	return err
}
