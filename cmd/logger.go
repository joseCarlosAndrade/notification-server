package main

// import (
// 	"os"

// 	"github.com/joseCarlosAndrade/notification-server/internal/core/config"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )





// /*
// logs:
// development:
// 2025-01-28T00:00:00.000+0800    DEBUG    development/main.go:7    This is a DEBUG message
// 2025-01-28T00:00:00.000+0800    INFO    development/main.go:8    This is an INFO message

// production: (debug messages are ignored)
// {"level":"info","ts":1737907200.0000000,"caller":"production/main.go:8","msg":"This is an INFO message"}
// {"level":"info","ts":1737907200.0000000,"caller":"production/main.go:9","msg":"This is an INFO message with fields","region":["us-west"],"id":2}


// logger.Info("this is a log info", zap.Int("int", 10))
// 	logger.Debug("this is a log info", zap.Int("int", 10))
// 	logger.Warn("this is a log info", zap.Int("int", 10))
// 	logger.Error("this is a log info", zap.Int("int", 10))

// 	zap.L().Info("global logger", zap.String("eheheh", "string value"))

// 	// childLogger := logger.With(zap.String("requestID", "10000")) // every log with childLogger will include these params
// 	// childLogger.Info("info", zap.String("key", "value"))

// */