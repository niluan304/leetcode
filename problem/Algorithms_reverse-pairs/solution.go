package main

import (
	"math"

	. "github.com/niluan304/leetcode/container/segtree"
)

// 两数之和 + 线段树
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func reversePairs(record []int) int {
	mn, mx := math.MaxInt, math.MinInt
	for _, v := range record {
		mn = min(mn, v)
		mx = max(mx, v)
	}

	tree := NewSumSegmentTree(mn, mx)

	ans := 0
	for _, v := range record {
		ans += tree.Query(v+1, math.MaxInt32)

		tree.Set(v, func(old int) int {
			return old + 1
		})
	}
	return ans
}
