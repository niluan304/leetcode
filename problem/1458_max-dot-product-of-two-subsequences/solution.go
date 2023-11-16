package main

import (
	"math"
)

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func maxDotProduct(nums1 []int, nums2 []int) int {
	var m, n = len(nums1), len(nums2)
	var cache = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = math.MinInt
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			if nums1[0] < 0 {
				return nums1[0] * _min(nums2[:j+1]...)
			}
			return nums1[0] * _max(nums2[:j+1]...)
		}
		if j == 0 {
			if nums2[0] < 0 {
				return nums2[0] * _min(nums1[:i+1]...)
			}
			return nums2[0] * _max(nums1[:i+1]...)
		}

		v := cache[i][j]
		if v > math.MinInt {
			return v
		}

		product := nums1[i] * nums2[j]
		v = _max(
			dfs(i-1, j),
			dfs(i, j-1),
			dfs(i-1, j-1)+_max(product, 0),
			product,
		)
		cache[i][j] = v
		return v
	}

	return dfs(m-1, n-1)
}

func _max(a ...int) int {
	x := a[0]
	for _, y := range a[1:] {
		if x < y {
			x = y
		}
	}
	return x
}

func _min(a ...int) int {
	x := a[0]
	for _, y := range a[1:] {
		if x > y {
			x = y
		}
	}
	return x
}

// dp 动态规划
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func maxDotProduct2(nums1 []int, nums2 []int) int {
	var m, n = len(nums1), len(nums2)
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = nums1[0] * nums2[0]
	for i := 1; i < n; i++ {
		dp[0][i] = Max2(dp[0][i-1], nums1[0]*nums2[i])
	}
	for i := 1; i < m; i++ {
		dp[i][0] = Max2(dp[i-1][0], nums1[i]*nums2[0])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			product := nums1[i] * nums2[j]
			dp[i][j] = Max2(
				Max2(
					dp[i-1][j],
					dp[i][j-1],
				),
				Max2(
					dp[i-1][j-1]+Max2(product, 0),
					product,
				),
			)
		}
	}

	return dp[m-1][n-1]
}

func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func maxDotProduct3(nums1 []int, nums2 []int) int {
	var m, n = len(nums1), len(nums2)
	var cache = make([][]*int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]*int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			if nums1[0] < 0 {
				return nums1[0] * _min(nums2[:j+1]...)
			}
			return nums1[0] * _max(nums2[:j+1]...)
		}
		if j == 0 {
			if nums2[0] < 0 {
				return nums2[0] * _min(nums1[:i+1]...)
			}
			return nums2[0] * _max(nums1[:i+1]...)
		}

		if cache[i][j] != nil {
			return *cache[i][j]
		}

		product := nums1[i] * nums2[j]
		v := _max(
			dfs(i-1, j),
			dfs(i, j-1),
			dfs(i-1, j-1)+_max(product, 0),
			product,
		)
		cache[i][j] = &v
		return v
	}

	return dfs(m-1, n-1)
}
