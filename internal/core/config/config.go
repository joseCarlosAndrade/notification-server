package config

const (
	DefaultAPIPort = "8080"
	AppTraceName   = "notification-server"
)

type AppInfo struct {
	Debug                        bool     `default:"true"`
	Development                  bool     `default:"true"`
	MongoURI                     string   `default:"mongodb://admin:password@localhost:27017"`
	MongoNotificationsDB         string   `default:"notifications"`
	MongoNotificationsCollection string   `default:"notifications"`
	RedpandaBrokers              []string `default:""`
	KafkaConsumerGroup           string   `default:""`
	NotificationTopic            string   `default:""`
	OtelExporterEndpoint         string   `default:""`     // not implemented yet
	UseCache                     bool     `default:"true"` // if true, uses redis as cache. if not, query everything everytime
	DefaultCacheTTLs             int      `default:"25"`   // default ttl in seconds for cache entries
}

var (
	App AppInfo
)
