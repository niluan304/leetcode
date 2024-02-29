package container

import "container/heap"

// MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap is a max-heap of ints.
type MaxHeap struct{ MinHeap }

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap.Less(j, i) }

// NewHeap
//
// 根据 data 和 less 初始化堆
//
// 当 x < y 时，
// 如果 less(x, y) = true, 返回最小堆；
// 如果 less(x, y) = false, 返回最大堆
func NewHeap[T any](data []T, less func(x, y T) bool) *Heap[T] {
	hp := &Heap[T]{
		data: data,
		less: less,
	}
	heap.Init(hp)
	return hp
}

// Heap implement generic heap
type Heap[T any] struct {
	data []T
	less func(x, y T) bool
}

func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}
