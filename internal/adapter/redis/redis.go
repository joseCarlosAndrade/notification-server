package redis

import (
	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
)

// Cache implements the port.Cache interface
type Cache struct {

}

// makes sure Cache implements the interface
var _ port.Cache = (*Cache)(nil)

func NewCache(ctx context.Context) Cache {
	return Cache{}
}

// Run implements port.Runner interface
func (s *Cache) Run(ctx context.Context) error {
	return nil
}

// Close implements port.Runner interface
func (s *Cache) Close(ctx context.Context) error {
	// todo: close connection and clean up

	return nil
}