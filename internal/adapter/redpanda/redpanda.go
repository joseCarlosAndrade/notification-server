package redpanda

import (
	"context"
	"errors"
	"time"

	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
)

// EventsHub implements the EventsHub interface
type EventsHub struct {
	service *port.Service
}

// makes sure EventsHub implements the interface
var _ port.EventsHub = (*EventsHub)(nil)

func NewEventsHub(ctx context.Context, serviceRepository *port.Service) EventsHub {
	return EventsHub{}
}

// Run implements port.Runner interface
func (s *EventsHub) Run(ctx context.Context) error {
	for {
		time.Sleep(time.Second*5) 
		return errors.New("new error from events hub")
	}
	return nil
}

// Close implements port.Runner interface
func (s *EventsHub) Close(ctx context.Context) error {
	// todo: close connection and clean up

	return nil
}