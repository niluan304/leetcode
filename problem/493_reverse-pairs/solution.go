package main

import (
	"math"

	. "github.com/niluan304/leetcode/container/segtree"
)

// 线段树
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func reversePairs(nums []int) int {
	mn, mx := math.MaxInt, math.MinInt
	for _, num := range nums {
		mn = min(mn, num)
		mx = max(mx, num)
	}

	tree := NewSumSegmentTree(mn, mx)

	ans := 0
	for _, num := range nums {
		ans += tree.Query(num*2+1, math.MaxInt)

		tree.Set(num, func(old int) int {
			return old + 1
		})
	}
	return ans
}
