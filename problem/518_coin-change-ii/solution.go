package main

// dp 动态规划
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func change(amount int, coins []int) int {
	// dp数组的含义：
	// dp[i][j]表示背包大小为j时，选到第i个物品，当前背包的最大总价值
	var n = len(coins)
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	// 初始化 dp数组
	dp[0][0] = 1
	for i, coin := range coins {
		for j := 0; j <= amount; j++ {
			// 递推公式：选与不选当前coin
			// 不选：dp[i-1][j]
			// 　选：dp[i][j-coin]
			// 题目要求至多, 因此 dp[i][j] = dp[i-1][j] + dp[i][j-coin]
			dp[i+1][j] = dp[i][j]
			if j >= coin {
				dp[i+1][j] += dp[i+1][j-coin]
			}
		}
	}

	//// debug: 查看dp数组
	//fmt.Println("amount", amount, "coins", coins)
	//for i := range dp {
	//	fmt.Println(dp[i])
	//}
	return dp[n][amount]
}

// dp 动态规划, 复用一维数组，以降低空间复杂度
// 时间复杂度：O(m*n)
// 空间复杂度：O(n)
func change2(amount int, coins []int) int {
	// dp数组的含义：
	// dp[i][j]表示背包大小为j时，选到第i个物品，当前背包的最大总价值
	var dp = make([]int, amount+1)

	// 初始化 dp数组
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			// 递推公式：选与不选当前coin
			// 不选：dp[i-1][j]
			// 　选：dp[i][j-coin]
			// 题目要求至多, 因此 dp[i][j] = dp[i-1][j] + dp[i][j-coin]
			dp[j] += dp[j-coin]
		}

		//// debug: 查看dp数组
		//fmt.Println("amount", amount, "coins", coins, "dp", dp)
	}

	return dp[amount]
}
