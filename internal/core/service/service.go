package service

import (
	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
)

// Service implements the port.Service interface
type Service struct {
	storage port.Storage
	cache port.Cache
}

func NewService(ctx context.Context, storageRepository port.Storage, cacheRepository port.Cache) Service {
	return Service{
		storage: storageRepository,
		cache: cacheRepository,
	}
}