package port

import "context"

type Runner interface {
	Run(ctx context.Context) error
	Close(ctx context.Context) error
}