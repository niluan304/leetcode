package main

import (
	. "github.com/niluan304/leetcode/container"
)

// 优先队列/堆排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func maxKelements(nums []int, k int) int64 {
	hp := NewMaxIntHeap(nums...)

	ans := 0
	for i := 0; i < k; i++ {
		top := hp.PopHead()
		hp.Insert((top + 3 - 1) / 3)
		ans += top
	}
	return int64(ans)
}

// 优先队列/堆排序
// 时间复杂度：$O(k \log n + n)$，其中 $n$ 是数组 $\textit{nums}$ 的长度。 构造优先队列需要的时间为 $O(n)$，每一轮操作需要的时间为 $O(\log n)$，一共有 $k$ 轮操作。
// 空间复杂度：O(1)，原地对数组建堆
func maxKelements2(nums []int, k int) int64 {
	hp := NewMaxIntHeap(nums...)

	ans := 0
	for i := 0; i < k; i++ {
		head := hp.Head()
		ans += head
		hp.Update(0, ceil(head, 3))

		// 等价于，先 heap.Pop 后 heap.Push
		// top := heap.Pop(&hp).(int)
		// ans += top
		// heap.Push(&hp, ceil(top, 3))
	}
	return int64(ans)
}

// 为了避免浮点数运算，我们可以用 $(x+3-1)/y$ 等价 $\lceil \dfrac{x}{3} \rceil$，其中 $/$ 表示整数除法。
func ceil(x, y int) int { return (x + y - 1) / y }
