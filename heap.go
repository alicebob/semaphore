package semaphore

import (
	"container/heap"
)

type item struct {
	elem     chan struct{}
	priority int
}

// Priority queue as per the example in container/heap.
//
// It's not goroutine safe, and don't pop() if there are no elements in the
// heap.
type priorityQueue []item

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}

func (pq *priorityQueue) Pop() interface{} {
	elem := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return elem
}

func (pq *priorityQueue) add(prio int, elem chan struct{}) {
	heap.Push(pq, item{
		elem:     elem,
		priority: prio,
	})
}

// pop() removes the priority item with the lowest value
func (pq *priorityQueue) pop() chan struct{} {
	return heap.Pop(pq).(item).elem
}
