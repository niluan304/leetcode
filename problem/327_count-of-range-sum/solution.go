package main

import (
	. "github.com/niluan304/leetcode/container/segtree"
	. "github.com/niluan304/leetcode/copypasta"
)

// 暴力穷举
// - 时间复杂度：O(n^2)。
// - 空间复杂度：O(1)。
// Deprecated: 超时
func bruteForce(nums []int, lower int, upper int) int {
	ans := 0
	for i := range nums {
		x := 0
		for _, num := range nums[i:] {
			x += num
			if lower <= x && x <= upper {
				ans++
			}
		}
	}
	return ans
}

// 两数之和 + 线段树
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	ans := 0

	// 将区间和转变为：「前缀和之差」
	preSum := make([]int, n+1)
	mn, mx := 0, 0 // 必须覆盖用于初始化的 set[0] = 1
	for i, num := range nums {
		sum := preSum[i] + num
		preSum[i+1] = sum

		mn = min(mn, sum)
		mx = max(mx, sum)

		//// 不初始化 set[0] = 1 的写法
		//if lower <= sum && sum <= upper {
		//	ans++
		//}
	}

	tree := NewSumSegmentTree(mn, mx)
	tree.Set(0, func(old int) int { return 1 }) // 初始化，用于匹配 preSum[i] 自身；也可以不初始化，在遍历时额外判断
	for _, x := range preSum[1:] {
		ans += tree.Query(x-upper, x-lower) // 两数之和「前缀和+哈希表」，已知 x，要求 lower <= x+y <= upper，求 y 的个数，
		tree.Set(x, func(old int) int {
			return old + 1
		})
	}
	return ans
}

// 前缀和 + 归并排序 + 逆序对
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
func countRangeSum2(nums []int, lower int, upper int) int {
	// 将区间和转变为：「前缀和之差」
	preSum := make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	ans := 0
	MergeSort(preSum, func(left, right []int) {
		// 统计下标对的数量
		l, r := 0, 0
		n := len(right)

		for _, v := range left {
			for l < n && right[l]-v < lower {
				l++
			}
			for r < n && right[r]-v <= upper {
				r++
			}
			ans += r - l
		}
	})
	return ans
}
