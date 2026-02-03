package port

import "context"

type EventsHub interface {
	Runner
	IsHealthy(ctx context.Context) error 
}