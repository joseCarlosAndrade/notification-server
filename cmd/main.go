package main

import (
	// "os"

	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/joseCarlosAndrade/notification-server/internal/core/config"
	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
	"github.com/joseCarlosAndrade/notification-server/internal/di"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)


func loadEnv() {
	_ = godotenv.Overload(".env")
	
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

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	container := di.NewContainer(ctx)

	if err := container.Run(ctx); err != nil  { // application somehow exited with failures
		log.L(context.Background()).Warn("container was closed due to errors")
	}

	log.L(context.Background()).Info("app exited")
}	