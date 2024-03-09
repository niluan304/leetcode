package main

import (
	"math"

	. "github.com/niluan304/leetcode/container/segtree"
	. "github.com/niluan304/leetcode/copypasta"
)

// 两数之和 + 线段树
// - 时间复杂度：O(nlogC)。C = mx-mn
// - 空间复杂度：O(n)。
func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	n := len(nums1)

	mn, mx := math.MaxInt32, math.MinInt32
	nums := make([]int, n)
	for i := range nums {
		nums[i] = nums1[i] - nums2[i]
		mn = min(mn, nums[i])
		mx = max(mx, nums[i])
	}

	ans := 0
	tree := NewSumSegmentTree(mn, mx)
	for _, num := range nums {
		ans += tree.Query(mn, num+diff)
		tree.Set(num, func(old int) int {
			return old + 1
		})
	}
	return int64(ans)
}

// 归并排序
//
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func numberOfPairs2(nums1 []int, nums2 []int, diff int) int64 {
	n := len(nums1)
	nums := make([]int, n)
	for i := range nums {
		nums[i] = nums1[i] - nums2[i]
	}

	ans := 0
	MergeSort(nums, func(left, right []int) {
		for i, j := 0, 0; j < len(right); j++ {
			for i < len(left) && left[i] <= right[j]+diff {
				i++
			}
			ans += i
		}
	})

	return int64(ans)
}
