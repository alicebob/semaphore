package semaphore

import (
	"context"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestSemaphoreSingle(t *testing.T) {
	var (
		res []int
		mu  sync.Mutex
		wg  sync.WaitGroup
		ctx = context.Background()
	)
	h := NewPriority(1)

	// fill that one slot
	h.Acquire(ctx, 99)

	// All of these will block.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(prio int) {
			h.Acquire(ctx, prio)
			mu.Lock()
			res = append(res, prio)
			mu.Unlock()
			h.Release()
			wg.Done()
		}(i)
	}
	time.Sleep(10 * time.Millisecond)

	// release the hounds
	h.Release()

	wg.Wait()
	if have, want := len(res), 100; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
	if !sort.IntsAreSorted(res) {
		t.Errorf("IntsAreSorted is false for: %v", res)
	}
}

func TestSemaphore(t *testing.T) {
	h := NewPriority(10)
	var (
		res []int
		mu  sync.Mutex
		wg  sync.WaitGroup
		ctx = context.Background()
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(prio int) {
			h.Acquire(ctx, -prio)
			mu.Lock()
			res = append(res, prio)
			mu.Unlock()
			h.Release()
			wg.Done()
		}(i)
	}

	wg.Wait()
	if have, want := len(res), 100; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}
