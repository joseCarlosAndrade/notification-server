package main

import (
	// "os"

	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/config"
	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
	"github.com/joseCarlosAndrade/notification-server/internal/di"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)




func loadEnv() {
	if err := envconfig.Process("app", &config.App); err != nil {
		panic(err)
	}
}



func main() {
	loadEnv()
	logger := log.InitLogger()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger) // replaces the global zap.L() logger with this one
	defer undo()

	ctx := context.Background()
	
	log.L(ctx).Info("starting app")
	
	di.NewContainer()
}	