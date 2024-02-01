package main

import (
	"cmp"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// dfs + 记忆化搜素
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func findTargetSumWays(nums []int, target int) int {
	var sum = Sum(nums) + target
	if sum%2 != 0 {
		return 0
	}

	var (
		half = sum / 2
		n    = len(nums)
	)
	var cache = make([]map[int]int, n)
	for i := range cache {
		cache[i] = make(map[int]int, half)
	}
	var dfs func(i int, target int) int
	dfs = func(i int, target int) int {
		if i < 0 {
			if target == 0 {
				return 1
			}
			return 0
		}

		value, ok := cache[i][target]
		if ok {
			return value
		}

		value = dfs(i-1, target)
		if target >= nums[i] {
			value += dfs(i-1, target-nums[i])
		}
		cache[i][target] = value
		return value
	}
	return dfs(n-1, half)
}

// 将 dfs+记忆化搜索 翻译为 递推/动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func findTargetSumWays2(nums []int, target int) int {
	var sum = target + Sum(nums)
	if sum%2 != 0 || sum < 0 {
		return 0
	}

	var (
		half = sum / 2
		n    = len(nums)
	)
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, half+1)
	}
	dp[0][0] = 1

	for i, v := range nums {
		//v := nums[i] // 物品的价值=其重量 v[i] = w[i]
		for j := 0; j <= half; j++ {
			dp[i+1][j] = dp[i][j]

			w := j - v // 背包剩余承受重量
			if w < 0 {
				continue
			}
			// 2. 确定递推公式/状态转移公式
			dp[i+1][j] += dp[i][w]
		}
	}

	//// 5. debug: 打印dp数组
	//for i := range dp {
	//	fmt.Println("half", half, "i", dp[i])
	//}
	return dp[n][half]
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

// 递推/动态规划 的空间优化
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func findTargetSumWays3(nums []int, target int) int {
	var sum = target + Sum(nums)
	if sum%2 != 0 || sum < 0 {
		return 0
	}

	var half = sum / 2
	var dp = make([]int, half+1)
	dp[0] = 1

	for _, v := range nums {
		//v := nums[i] // 物品的价值=其重量 v[i] = w[i]
		for j := half; j >= 0; j-- {
			dp[j] = dp[j]

			w := j - v // 背包剩余承受重量
			if w < 0 {
				continue
			}
			// 2. 确定递推公式/状态转移公式
			dp[j] += dp[w]
		}

		//// 5. debug: 打印dp数组
		//fmt.Println("half", half, "i", dp)
	}
	return dp[half]
}

func findTargetSumWays4(nums []int, target int) int {
	sum := Sum(nums) + target
	return ZeroOneWaysToSum(nums, sum/2)
}
