package main

import (
	"container/heap"
)

// 优先队列/堆排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func findKthLargest(nums []int, k int) int {
	var h = new(Heap)
	for _, num := range nums {
		heap.Push(h, num)
	}

	var ans = 0
	for i := 0; i < k; i++ {
		ans = heap.Pop(h).(int)
	}
	return ans
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] } // 大根堆
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
