package di

import (
	"go.uber.org/zap"
)

func NewContainer() {
	zap.L().Info("this is from the global logger")
}