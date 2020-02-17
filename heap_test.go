package semaphore

import (
	"testing"
)

func TestHeap(t *testing.T) {
	i1 := make(chan struct{})
	i2 := make(chan struct{})
	i3 := make(chan struct{})

	h := &priorityQueue{}
	h.Add(4, i1)
	h.Add(2, i2)
	h.Add(8, i3)

	if have, want := h.Pop(), i2; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := h.Pop(), i1; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := h.Pop(), i3; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
