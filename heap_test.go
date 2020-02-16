package semaphore

import (
	"testing"
)

func TestHeap(t *testing.T) {
	i1 := make(chan struct{})
	i2 := make(chan struct{})
	i3 := make(chan struct{})

	h := &priorityQueue{}
	h.add(4, i1)
	h.add(2, i2)
	h.add(8, i3)

	if have, want := h.pop(), i2; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := h.pop(), i1; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := h.pop(), i3; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
