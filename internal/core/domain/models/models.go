package models

import "time"

// NotificationRecord represents the record received by the eventshub layer
type NotificationRecord struct {
	Service string    `json:"service"` // Service represents the service that produced this notification
	Message string    `json:"message"` // Message is the data
	SentAt  time.Time `json:"sentAt"`
}

type Notification struct {
	ID      string     `json:"_id"`
	Service string     `json:"service"`
	Message string     `json:"message"`
	IsRead  bool       `json:"isRead"`
	SentAt  time.Time  `json:"sentAt"`
	ReadAt  *time.Time `json:"readAt"` // might not exist yet
}

// LastTime represnets the filter for getting notifications from the last day-hour-minute
type LastTime struct {
	Days int
	Hours int
	Minutes int
}