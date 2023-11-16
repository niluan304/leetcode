package main

import (
	"math"
	"sort"
)

// dp, 动态规划, 本质为爬楼梯
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func minCapability(nums []int, k int) int {
	var n = len(nums)
	var DP = func(capability int) bool {
		var dp = make([]int, n+2)
		for i, num := range nums {
			dp[i+2] = dp[i+1]
			if capability >= num {
				dp[i+2] = _max(dp[i+1], dp[i]+1)
			}
		}
		return dp[n+1] >= k
	}

	var ans = math.MaxInt
	for _, num := range nums {
		if num > ans {
			continue
		}
		if DP(num) {
			ans = _min(ans, num)
		}
	}
	return ans
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

// dp, 动态规划, 本质为爬楼梯
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func minCapability2(nums []int, k int) int {
	var DP = func(capability int) bool {
		var dp0, dp1 = 0, 0
		for _, num := range nums {
			dp2 := dp1
			if capability >= num {
				dp2 = _max(dp2, dp0+1)
			}
			dp1, dp0 = dp2, dp1
		}
		return dp1 >= k
	}

	var ans = math.MaxInt
	for _, num := range nums {
		if num > ans {
			continue
		}
		if DP(num) {
			ans = _min(ans, num)
		}
	}
	return ans
}

// dp, 动态规划, 本质为爬楼梯
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func minCapability3(nums []int, k int) int {
	var DP = func(capability int) bool {
		var dp0, dp1 = 0, 0
		for _, num := range nums {
			dp2 := dp1
			if capability >= num {
				dp2 = _max(dp2, dp0+1)
			}
			dp1, dp0 = dp2, dp1
		}
		return dp1 >= k
	}

	return sort.Search(1e9, DP)
}
