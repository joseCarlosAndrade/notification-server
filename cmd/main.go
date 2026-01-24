package main

import (
	"os"

	"github.com/joseCarlosAndrade/notification-server/internal/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
logs:
development:
2025-01-28T00:00:00.000+0800    DEBUG    development/main.go:7    This is a DEBUG message
2025-01-28T00:00:00.000+0800    INFO    development/main.go:8    This is an INFO message

production: (debug messages are ignored)
{"level":"info","ts":1737907200.0000000,"caller":"production/main.go:8","msg":"This is an INFO message"}
{"level":"info","ts":1737907200.0000000,"caller":"production/main.go:9","msg":"This is an INFO message with fields","region":["us-west"],"id":2}
*/

var logger *zap.Logger

func loadEnv() {

}

func initLoger() {
	var encoderConfig zapcore.EncoderConfig
	
	if config.App.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig() // plain text encoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig() // JSON encoder format TODO: THIS IS NOT WORKING
	}
	
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // time format
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // display different colors for different levels
	// encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // display full file path (/home/username/.../file.go)

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller()) // this add caller also display the path to the file, but relative

}

func main() {
	loadEnv()
	initLoger()
	defer logger.Sync()

	logger.Info("this is a log info", zap.Int("int", 10))
	logger.Debug("this is a log info", zap.Int("int", 10))
	logger.Warn("this is a log info", zap.Int("int", 10))
	logger.Error("this is a log info", zap.Int("int", 10))

}	