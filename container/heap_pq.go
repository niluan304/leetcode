package container

import "container/heap"

// An PriorityQueueNode is something we manage in a priority queue.
type PriorityQueueNode[T any] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by Update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (o *PriorityQueueNode[T]) Value() T {
	return o.value
}

func (o *PriorityQueueNode[T]) Priority() int {
	return o.priority
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*PriorityQueueNode[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityQueueNode[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Insert 入队操作，执行 heap.Push(pq, &PriorityQueueNode{value, priority})
func (pq *PriorityQueue[T]) Insert(value T, priority int) {
	heap.Push(pq, &PriorityQueueNode[T]{value: value, priority: priority})
}

// Remove 出队操作 执行 heap.Pop(pq).(*PriorityQueueNode[T])
func (pq *PriorityQueue[T]) Remove() *PriorityQueueNode[T] {
	return heap.Pop(pq).(*PriorityQueueNode[T])
}

// Update modifies the priority and value of an PriorityQueueNode in the queue.
func (pq *PriorityQueue[T]) Update(item *PriorityQueueNode[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
