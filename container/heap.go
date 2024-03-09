package container

import (
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
	h := &Heap[T]{
		data: data,
		less: less,
	}

	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

// NewEmptyHeap
//
// 根据 less 返回最小堆/最大堆
//
// 当 x < y 时，
// 如果 less(x, y) = true, 返回最小堆；
// 如果 less(x, y) = false, 返回最大堆
func NewEmptyHeap[T any](less func(x, y T) bool) *Heap[T] {
	h := &Heap[T]{
		data: nil,
		less: less,
	}
	return h
}

// Heap implement generic heap
type Heap[T any] struct {
	data []T
	less func(x, y T) bool
}

func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

// copy form standard library "container/heap"
func (h *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// copy form standard library "container/heap"
func (h *Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

// Peek 仅返回堆的根节点的值，不作删除
// 非空的堆才可以调用，空堆调用会触发 panic(slices out of size)
func (h *Heap[T]) Peek() T {
	return h.data[0]
}

// Push 入队操作，将 value 插入 Heap 中
func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.up(h.Len() - 1)
}

// Pop 出队操作，弹出并返回 Heap 的根，等价于 Remove(0)
func (h *Heap[T]) Pop() T {
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)

	old := h.data
	x := old[n]
	h.data = old[0:n]
	return x
}

// Update 更新操作，将 Heap 中的节点修改为给定的 value
func (h *Heap[T]) Update(i int, value T) {
	h.data[i] = value

	if !h.down(i, h.Len()) {
		h.up(i)
	}
}

// Remove 删除操作，从 Heap 中移除指定节点，并返回
func (h *Heap[T]) Remove(i int) T {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.Pop()
}

func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}

// Values 返回堆的值的拷贝
func (h *Heap[T]) Values() []T {
	return slices.Clone(h.data)
}
