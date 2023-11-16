package main

import (
	"sort"
)

// dp 动态规划 LIS
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var n = len(obstacles)
	var dp = make([]int, n)

	for i := 0; i < n; i++ {
		length := 0

		for j := 0; j < i; j++ {
			if obstacles[i] < obstacles[j] {
				continue
			}
			length = _max(length, dp[j])
		}
		dp[i] = length + 1
	}

	return dp
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 贪心 + 二分
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func longestObstacleCourseAtEachPosition2(obstacles []int) []int {
	var g []int
	var ans = make([]int, len(obstacles))
	for i, x := range obstacles {
		j := sort.SearchInts(g, x+1)
		if j == len(g) { // >=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}

		ans[i] = j + 1
	}
	return ans
}

// 方法二：贪心 + 二分查找
// 前置知识：二分查找
// 见【基础算法精讲 04】。
//
// 写法一：额外空间
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/longest-increasing-subsequence/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func lengthOfLIS(nums []int) []int {
	var g []int

	var ans = make([]int, len(nums))
	for i, x := range nums {
		j := sort.SearchInts(g, x+1)
		if j == len(g) { // >=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}

		ans[i] = j + 1
	}
	return ans
}
