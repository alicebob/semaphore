// Package semaphore is a semaphore with a priority queue backend.
// It is a stripped down /x/sync/semaphore with a container/heap queue.
package semaphore

import (
	"context"
	"sync"
)

// Priority is a semaphore with priority waiters. Lowest priority value gets woken up first.
type Priority struct {
	size    int
	cur     int
	mu      sync.Mutex
	waiters priorityQueue
}

// NewPriority makes a new priority semaphore with the given size. Size must be at least 1.
func NewPriority(size int) *Priority {
	if size < 1 {
		panic("priority semaphore size must be at least 1")
	}
	return &Priority{
		size: size,
	}
}

// Acquire a slot in the semaphore. Every call to Acquire() needs to have a
// corresponding call to Release().
// Lower prio numbers will get unblocked first. If there are multiple
// Acquire()s with the same prio it is not defined which one goes first.
func (p *Priority) Acquire(ctx context.Context, prio int) error {
	p.mu.Lock()
	if p.size-p.cur >= 1 { // && p.waiters.Len() == 0 {
		p.cur++
		p.mu.Unlock()
		return nil
	}

	ready := make(chan struct{})
	p.waiters.add(prio, ready)
	p.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ready:
		return nil
	}
}

// Release() a previous Acquire()
func (p *Priority) Release() {
	p.mu.Lock()
	p.cur--
	if p.cur < 0 {
		p.mu.Unlock()
		panic("semaphore: released more than held")
	}
	if p.waiters.Len() > 0 {
		p.cur++
		close(p.waiters.pop())
	}
	p.mu.Unlock()
}
