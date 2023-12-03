package queue

import "context"

type IQueue interface {
	Push(ctx context.Context, data any) (int, error)
	Fetch(ctx context.Context, offset int) (any, error)
	Clear() error
	Len() int
}
