package main

import (
	"math"

	. "github.com/niluan304/leetcode/container/segtree"
	. "github.com/niluan304/leetcode/copypasta"
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

// 归并排序 + 逆序对
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func reversePairs2(nums []int) int {
	ans := 0
	MergeSort(nums, func(left, right []int) {
		// 统计重要翻转对 (i,j) 的数量
		// 由于 left 和 right 均为有序，可以用两个指针同时遍历
		for i, j := 0, 0; i < len(left); i++ {
			for j < len(right) && left[i] > 2*right[j] {
				j++
			}
			ans += j
		}
	})
	return ans
}
