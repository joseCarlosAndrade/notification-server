package models


type NotificationRecord struct {
	Service string `json:"service"` // Service represents the service that produced this notification
	Message string `json:"message"` // Message is the data 
}