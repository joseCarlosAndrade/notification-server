package port

import "context"

type Storage interface {
	Runner
	IsHealthy(ctx context.Context) error 
}