package mongo

import (
	"context"
	"fmt"
	"time"

	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
)

// Storage implements the port.Storage
type Storage struct {
	client         *mongo.Client
	dbName         string
	collectionName string
}

var _ port.Storage = (*Storage)(nil) // ensures Storage implements port.Storage

func NewStorage(ctx context.Context, connectionStr, mongoDB, mongoCollection string) (Storage, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionStr))
	if err != nil {
		log.L(ctx).Error("invalid mongo config", zap.Error(err))
		return Storage{}, err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.L(ctx).Error("database not connected. ping failed", zap.Error(err))
		return Storage{}, fmt.Errorf("could not connect to database: %w", err)
	}

	log.L(ctx).Info("successfully connected to mongodb")

	return Storage{
		client: client,
		dbName: mongoDB,
		collectionName: mongoCollection,
	}, nil
}

// Run implements port.Runner interface
func (s *Storage) Run(ctx context.Context) error {
	return nil
}

// Close implements port.Runner interface
func (s *Storage) Close(ctx context.Context) error {
	// todo: close connection and clean up

	return s.client.Disconnect(ctx)
}

func (s *Storage) IsHealthy(ctx context.Context) error {
	return s.client.Ping(ctx, readpref.Primary())
}
