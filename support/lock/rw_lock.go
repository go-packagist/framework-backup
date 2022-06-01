package lock

import "sync"

// RwLock is a read-write mutex.
type RwLock struct {
	rw sync.RWMutex
}

// NewRwLock creates a new read-write lock.
func NewRwLock() *RwLock {
	return &RwLock{
		rw: sync.RWMutex{},
	}
}

// Run runs the given function in a read-write lock.
func (l *RwLock) Run(fn func()) {
	l.rw.Lock()
	defer l.rw.Unlock()

	fn()
}
