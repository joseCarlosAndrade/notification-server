package port

import (
	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/models"
)

type Service interface {
	// todo: define service operations that saves noticiation: tries cache first, then store it on mongo
	
	// SaveNewNotification generates an id, stores notification in db and in cache (if available)
	SaveNewNotification(ctx context.Context, notification *models.NotificationRecord) error

}