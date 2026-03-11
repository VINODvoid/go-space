package kvstore

import (
	"sync"
	"time"
)

type entry[V any] struct {
	value  V
	expiry time.Time
}
type store[K comparable, V any] struct {
	data          map[K]entry[V]
	mu            sync.RWMutex
	sweepInterval time.Duration
}
