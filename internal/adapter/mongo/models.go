package mongo

import "time"

// Notification is the document schema for mongodb notification storage
type Notification struct {
	ID      string     `bson:"_id"`
	Service string     `bson:"service"`
	Message string     `bson:"message"`
	IsRead  bool       `bson:"isRead"`
	SentAt  time.Time  `bson:"sentAt"`
	ReadAt  *time.Time `bson:"readAt"` // might not exist yet
}
