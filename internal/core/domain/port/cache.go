package port

import "context"

type Cache interface {
	Runner
	IsHealthy(ctx context.Context) error 
}