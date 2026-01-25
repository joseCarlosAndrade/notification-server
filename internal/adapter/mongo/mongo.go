package mongo

import (
	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
)

// Storage implements the port.Storage 
type Storage struct {

}

var _ port.Storage = (*Storage)(nil) // ensures Storage implements port.Storage

func NewStorage(ctx context.Context) Storage {
	return Storage{}
}

// Run implements port.Runner interface
func (s *Storage) Run(ctx context.Context) error {
	return nil
}

// Close implements port.Runner interface
func (s *Storage) Close(ctx context.Context) error {
	// todo: close connection and clean up

	return nil
}