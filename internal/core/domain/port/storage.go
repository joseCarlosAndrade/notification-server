package port

import (
	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/models"
)

type Storage interface {
	Runner
	IsHealthy(ctx context.Context) error 

	StoreNewNotification(ctx context.Context, notification *models.NotificationRecord, id string) error 
	MarkNotificationAsRead(ctx context.Context, notificationID string) error
	GetAllNotificationsByTime(ctx context.Context, serviceName string, filter models.LastTime) ([]*models.NotificationRecord, error)
	GetNonReadNotifications(ctx context.Context, serviceName string) ([]*models.NotificationRecord, error)
}