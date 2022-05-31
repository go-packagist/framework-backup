package lock

import "sync"

type RWLock struct {
	rw sync.RWMutex
}

func NewRWLock() *RWLock {
	return &RWLock{
		rw: sync.RWMutex{},
	}
}

func (l *RWLock) Run(fn func()) {
	l.rw.Lock()
	defer l.rw.Unlock()

	fn()
}
