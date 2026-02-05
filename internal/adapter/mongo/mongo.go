package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain"
	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/models"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	notificationCollection *mongo.Collection
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
		notificationCollection: client.Database(mongoDB).Collection(mongoCollection),
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

func (s *Storage) StoreNewNotification(ctx context.Context, notification *models.NotificationRecord, id string) error {
	// start span here

	mongoNotification := transformNotificationToMongo(notification, id)

	filter := bson.M{"_id" : id}

	update := bson.M{
		"$set" : *mongoNotification,
	}

	opts := options.UpdateOne().SetUpsert(true)

	// using upsert to avoid duplicate if the service tries to save the same notification (idempotency)
	_, err := s.notificationCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.L(ctx).Error("could not insert new notification", 
			zap.String("id", mongoNotification.ID),
			zap.String("service",mongoNotification.Service),
			zap.Error(err))
		
			return fmt.Errorf("failed to insert in mongodb: %w", err)
	}

	log.L(ctx).Debug("successfully stored new notification in mongo",
		zap.String("id", mongoNotification.ID),
		zap.String("service",mongoNotification.Service))

	return nil
}

func (s *Storage)MarkNotificationAsRead(ctx context.Context, notificationID string) error {
	if notificationID == "" {
		return fmt.Errorf("notificationID canont be empty")
	}	

	filter := bson.M{
		"_id" : notificationID,
	}

	now := domain.NewNowTimeString()

	update := bson.M{
		"$set" : bson.M{
			"readAt" : now,
			"isRead" : true,
		},
	}

	res, err := s.notificationCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.L(ctx).Error("could not update document in mongodb",
		zap.String("id", notificationID),
		zap.Error(err))

		return fmt.Errorf("could not update document in mongodb: %w", err)
	}

	if res.MatchedCount == 0 {
		log.L(ctx).Error("no documents matched this filter", zap.String("id", notificationID))
		return fmt.Errorf("no documents matched this filter")
	}

	log.L(ctx).Info("notification successfully marked as read", zap.String("id", notificationID))
	
	return nil
}

func transformNotificationToMongo(notification *models.NotificationRecord, id string) *Notification {
	return &Notification{
		ID: id,
		Service: notification.Service,
		Message: notification.Message,
		IsRead: false,
		SentAt: notification.SentAt,
	}
}

func (s *Storage) GetAllNotificationsByTime(ctx context.Context, serviceName string, filter models.LastTime) ([]*models.NotificationRecord, error) {
	return nil, nil
}
func (s *Storage) GetNonReadNotifications(ctx context.Context, serviceName string) ([]*models.NotificationRecord, error) {
	return nil, nil
}