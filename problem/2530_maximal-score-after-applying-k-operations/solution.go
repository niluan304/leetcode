package main

import (
	"container/heap"
)

// 优先队列/堆排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func maxKelements(nums []int, k int) int64 {
	var hp = new(IntHeap)
	for _, num := range nums {
		heap.Push(hp, num)
	}

	var ans = 0
	for i := 0; i < k; i++ {
		top := heap.Pop(hp).(int)
		ans += top
		heap.Push(hp, (top+3-1)/3)
	}
	return int64(ans)
}

// 优先队列/堆排序
// 时间复杂度：$O(k \log n + n)$，其中 $n$ 是数组 $\textit{nums}$ 的长度。 构造优先队列需要的时间为 $O(n)$，每一轮操作需要的时间为 $O(\log n)$，一共有 $k$ 轮操作。
// 空间复杂度：O(1)，原地对数组建堆
func maxKelements2(nums []int, k int) int64 {
	var hp IntHeap = nums
	heap.Init(&hp)

	var ans = 0
	for i := 0; i < k; i++ {
		ans += hp[0]
		hp[0] = ceil(hp[0], 3)
		heap.Fix(&hp, 0)

		// 等价于，先 heap.Pop 后 heap.Push
		//top := heap.Pop(&hp).(int)
		//ans += top
		//heap.Push(&hp, ceil(top, 3))
	}
	return int64(ans)
}

// 为了避免浮点数运算，我们可以用 $(x+3-1)/y$ 等价 $\lceil \dfrac{x}{3} \rceil$，其中 $/$ 表示整数除法。
func ceil(x, y int) int { return (x + y - 1) / y }

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // 大根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
