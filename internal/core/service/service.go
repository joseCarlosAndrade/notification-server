package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/models"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
	"go.uber.org/zap"
)

// Service implements the port.Service interface
type Service struct {
	storage port.Storage
	cache   port.Cache
}

func NewService(ctx context.Context, storageRepository port.Storage, cacheRepository port.Cache) Service {
	return Service{
		storage: storageRepository,
		cache:   cacheRepository,
	}
}

func (s *Service) SaveNewNotification(ctx context.Context, notification *models.NotificationRecord) error {
	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("could not generate id: %w", err)
	}

	err = s.storage.StoreNewNotification(ctx, notification, id.String())
	if err != nil {
		log.L(ctx).Error("could not store new notification",
			zap.String("id", id.String()),
			zap.Error(err))

		return fmt.Errorf("could not store new notification: %w", err)
	}

	// todo: store in cache

	log.L(ctx).Info("notification successfully stored",
		zap.String("id", id.String()))


	return nil
}
