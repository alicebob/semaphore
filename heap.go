package semaphore

import (
	"container/heap"
)

// item is what's stored in the priority queue
type item struct {
	elem     chan struct{}
	priority int
}

// Priority queue as per the example in container/heap.
//
// It's not goroutine safe, and don't pop() if there are no elements in the
// heap.
type priorityQueue struct {
	heap priorityHeap
}

// Add() adds an element to the priority queue
func (pq *priorityQueue) Add(prio int, elem chan struct{}) {
	heap.Push(&pq.heap, item{
		elem:     elem,
		priority: prio,
	})
}

// Pop() removes the priority item with the lowest value
func (pq *priorityQueue) Pop() chan struct{} {
	return heap.Pop(&pq.heap).(item).elem
}

// Number of elements in the queue
func (pq *priorityQueue) Len() int {
	return len(pq.heap)
}

// priorityHeap implements the Interface for container/heap
type priorityHeap []item

func (pq priorityHeap) Len() int {
	return len(pq)
}

func (pq priorityHeap) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityHeap) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}

func (pq *priorityHeap) Pop() interface{} {
	elem := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return elem
}
