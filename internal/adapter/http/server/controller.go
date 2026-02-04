package server

import (
	"context"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
)

// Controller implements the port.Controller interface
type Controller struct {
	service *port.Service
}

// makes sure Controller implements the interface
var _ port.Controller = (*Controller)(nil)

func NewController(ctx context.Context, serviceRepository *port.Service) Controller {
	return Controller{}
}

// Run implements port.Runner interface
func (s *Controller) Run(ctx context.Context) error {
	for {}
	return nil
}

// Close implements port.Runner interface
func (s *Controller) Close(ctx context.Context) error {
	// todo: close connection and clean up

	return nil
}

func (s *Controller) IsHealthy(ctx context.Context) error  {
	return nil
}