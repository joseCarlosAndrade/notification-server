package port

import "context"

type Controller interface {
	Runner	
	IsHealthy(ctx context.Context) error 
}