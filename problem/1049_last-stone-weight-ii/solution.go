package main

import (
	"math"
)

// dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func lastStoneWeightII(stones []int) int {
	var sum = Sum(stones)
	var target = sum / 2
	var n = len(stones)

	// 1. 确定dp(dp table)数组以及下标的含义
	// 01背包, dp[i][w], dp[取到第i件物品,背包承重]=背包内最大重量
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	// 3. dp数组如何初始化
	// dp[i][0] = 0
	// dp[0][j] = 0
	for i, v := range stones {
		//v := nums[i] // 物品的价值=其重量 v[i] = w[i]
		for j := 1; j <= target; j++ {
			dp[i+1][j] = dp[i][j]

			w := j - v // 背包剩余承受重量
			if w < 0 {
				continue
			}
			// 2. 确定递推公式/状态转移公式
			dp[i+1][j] = _max(dp[i+1][j], v+dp[i][w])
		}
	}
	//// 5. debug: 打印dp数组
	//for i := range dp {
	//	fmt.Println("target", target, "i", dp[i])
	//}

	return sum - dp[n][target]*2
}

func Sum(list []int) int {
	var sum = 0
	for _, n := range list {
		sum += n
	}

	return sum
}

func _max(list ...int) int {
	var ans = math.MinInt
	for _, n := range list {
		if ans < n {
			ans = n
		}
	}

	return ans
}

// dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n*2)
func lastStoneWeightII2(stones []int) int {
	var sum = Sum(stones)
	var target = sum / 2

	// 1. 确定dp(dp table)数组以及下标的含义
	// 01背包, dp[i][w], dp[取到第i件物品,背包承重]=背包内最大重量
	var dp = make([]int, target+1)

	// 3. dp数组如何初始化
	// dp[0] = 0
	for _, v := range stones {
		//v := nums[i] // 物品的价值=其重量 v[i] = w[i]
		var dp2 = append([]int{}, dp...)
		for j := 1; j <= target; j++ {
			w := j - v // 背包剩余承受重量
			if w < 0 {
				continue
			}
			// 2. 确定递推公式/状态转移公式
			dp2[j] = _max(dp2[j], v+dp[w])
		}
		dp = dp2
		//// 5. debug: 打印dp数组
		//fmt.Println("target", target, "dp", dp)

	}
	return sum - dp[target]*2
}
