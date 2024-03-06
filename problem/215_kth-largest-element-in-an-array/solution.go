package main

import (
	. "github.com/niluan304/leetcode/container"
)

// 优先队列/堆排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func findKthLargest(nums []int, k int) int {
	h := NewMaxIntHeap(nums...)

	ans := 0
	for i := 0; i < k; i++ {
		ans = h.PopHead()
	}
	return ans
}
