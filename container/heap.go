package container

import (
	"container/heap"
	"slices"
)

// NewMinIntHeap return a min-heap of ints.
func NewMinIntHeap(data ...int) *Heap[int] {
	return NewHeap[int](data, func(x, y int) bool {
		return x < y
	})
}

// NewMaxIntHeap return a max-heap of ints.
func NewMaxIntHeap(data ...int) *Heap[int] {
	return NewHeap[int](data, func(x, y int) bool {
		return x > y
	})
}

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

// NewEmptyHeap
//
// 根据 less 返回最小堆/最大堆
//
// 当 x < y 时，
// 如果 less(x, y) = true, 返回最小堆；
// 如果 less(x, y) = false, 返回最大堆
func NewEmptyHeap[T any](less func(x, y T) bool) *Heap[T] {
	hp := &Heap[T]{
		data: nil,
		less: less,
	}
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

// Heap 的操作
// - Insert
// - PopHead
// - Update
// - Remove

// Insert 入队操作，将 value 插入 Heap 中
// 简写：heap.Push(h, value)
func (h *Heap[T]) Insert(value T) {
	heap.Push(h, value)
}

// PopHead 出队操作，弹出并返回 Heap 的根
// 简写：heap.Pop(h).(T)
func (h *Heap[T]) PopHead() T {
	return heap.Pop(h).(T)
}

// Update 更新操作，将 Heap 中的节点修改为给定的 value
// 简写：h.data[i] = value; heap.Fix(h, i)
func (h *Heap[T]) Update(i int, value T) {
	h.data[i] = value
	heap.Fix(h, i)
}

// Remove 删除操作，从 Heap 中移除指定节点
// 简写：heap.Remove(h, i)
func (h *Heap[T]) Remove(i int) {
	heap.Remove(h, i)
}

// Head 返回堆的根节点的值
func (h *Heap[T]) Head() T {
	return h.data[0]
}

// Data 返回堆的值的拷贝
func (h *Heap[T]) Data() []T {
	return slices.Clone(h.data)
}
