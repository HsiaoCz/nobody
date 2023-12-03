package queue

import "sync"

type MQueue struct {
	// Concurrency lock
	mc sync.Mutex
	// len(data)
	len int
	// data
	data []any
	// MaxLen
	maxsize int
}

func NewQueue(size int) *MQueue {
	return &MQueue{
		data:    make([]any, 0),
		mc:      sync.Mutex{},
		len:     0,
		maxsize: size,
	}
}

