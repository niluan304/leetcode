package main

import (
	"math"
)

// 1. 确定dp(dp table)数组以及下标的含义
// 2. 确定递推公式/状态转移公式
// 3. dp数组如何初始化
// 4. 确定遍历顺序
// 5. debug: 打印dp数组

func canPartition(nums []int) bool {
	var n = len(nums)

	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	// 1. 确定dp(dp table)数组以及下标的含义
	// 01背包, dp[i][w], dp[取到第i件物品,背包承重]=背包内最大重量
	total := sum / 2
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, total+1)
	}
	// 3. dp数组如何初始化
	// dp[i][0] = 0
	// dp[0][j] = 0

	// 4. 确定遍历顺序
	// 先遍历物品, 后遍历背包重量
	for i, num := range nums {
		//v := nums[i] // 物品的价值=其重量 v[i] = w[i]
		for j := 1; j <= total; j++ {
			dp[i+1][j] = dp[i][j]

			w := j - num // 背包剩余承受重量
			if w < 0 {
				continue
			}
			// 2. 确定递推公式/状态转移公式
			dp[i+1][j] = _max(dp[i+1][j], num+dp[i][w])
		}
	}
	//// 5. debug: 打印dp数组
	//for i := range dp {
	//	fmt.Println("total", total, "i", dp[i])
	//}
	return dp[n][total] == total
}

func _max(list ...int) int {
	var ans = math.MinInt
	for _, item := range list {
		if ans < item {
			ans = item
		}
	}

	return ans
}

// dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n*2)
func canPartition2(nums []int) bool {
	var n = len(nums)

	if n < 2 {
		return false
	}

	sum, maxNum := 0, math.MinInt
	for _, num := range nums {
		sum += num

		if maxNum < num {
			maxNum = num
		}
	}

	if sum%2 == 1 {
		return false
	}

	// 1. 确定dp(dp table)数组以及下标的含义
	// 01背包, dp[i][w], dp[取到第i件物品,背包承重]=背包内最大重量
	target := sum / 2
	if target < maxNum {
		return false
	}
	if target == maxNum {
		return true
	}

	var dp = make([]int, target+1)

	// 3. dp数组如何初始化
	// dp[0] = 0

	// 4. 确定遍历顺序
	// 先遍历物品, 后遍历背包重量
	for _, v := range nums {
		//v := nums[i] // 物品的价值=其重量 v[i] = w[i]
		for j := target; j >= v; j-- {
			// 2. 确定递推公式/状态转移公式
			dp[j] = _max(dp[j], v+dp[j-v])
		}
	}
	//// 5. debug: 打印dp数组
	//for i := range dp {
	//	fmt.Println("target", target, "i", i, " dp[i]", dp[i])
	//}
	return dp[target] == target
}

// dp
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func canPartition3(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, maxNum := 0, math.MinInt
	for _, v := range nums {
		sum += v
		if maxNum < v {
			maxNum = v
		}
	}

	target := sum / 2
	if maxNum > target || target*2 != sum {
		return false
	}
	if maxNum == target {
		return true
	}

	// 1. 确定dp(dp table)数组以及下标的含义
	dp := make([]bool, target+1)

	//// 5. debug: 打印dp数组
	//fmt.Printf("\n\t")
	//for x := range dp {
	//	fmt.Printf("%d\t", x)
	//}
	//fmt.Println()

	// 3. dp数组如何初始化
	dp[0] = true
	// 4. 确定遍历顺序
	for _, v := range nums {
		for j := target; j >= v; j-- {
			// 2. 确定递推公式/状态转移公式
			dp[j] = dp[j] || dp[j-v]
		}

		//// 5. debug: 打印dp数组
		//fmt.Printf("%d\t", v)
		//for _, b := range dp {
		//	var x = 0
		//	if b {
		//		x = 1
		//	}
		//	fmt.Printf("%d\t", x)
		//}
		//fmt.Println()
	}
	return dp[target]
}
