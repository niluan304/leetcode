package main

import (
	"math"

	. "github.com/niluan304/leetcode/container/segtree"
	. "github.com/niluan304/leetcode/copypasta"
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

// 归并排序
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func reversePairs2(record []int) int {
	ans := 0

	MergeSort(record, func(left, right []int) {
		for i, j := 0, 0; i < len(left); i++ {
			for j < len(right) && left[i] > right[j] {
				j++
			}
			ans += j
		}
	})
	return ans
}
