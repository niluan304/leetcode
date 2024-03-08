package main

import (
	. "github.com/niluan304/leetcode/container/segtree"
)

// 两数之和 + 线段树
// - 时间复杂度：O(nlogC)。C = 1e4 - (-1e4)
// - 空间复杂度：O(n)。
func countSmaller(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	tree := NewSumSegmentTree(-1e4, 1e4)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		ans[i] = tree.Query(-1e4, num-1)

		tree.Set(num, func(old int) int {
			return old + 1
		})
	}
	return ans
}
