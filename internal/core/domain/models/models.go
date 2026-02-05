package models

import "time"

// NotificationRecord represents the record received by the eventshub layer
type NotificationRecord struct {
	Service string    `json:"service"` // Service represents the service that produced this notification
	Message string    `json:"message"` // Message is the data
	SentAt  time.Time `json:"sentAt"`
}

// LastTime represnets the filter for getting notifications from the last day-hour-minute
type LastTime struct {
	Days int
	Hours int
	Minutes int
}