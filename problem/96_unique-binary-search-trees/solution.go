package main

// 1. 确定dp(dp table)数组以及下标的含义
// 2. 确定递推公式/状态转移公式
// 3. dp数组如何初始化
// 4. 确定遍历顺序
// 5. debug: 打印dp数组

func numTrees(n int) int {
	// 1. 确定dp(dp table)数组以及下标的含义
	var dp = make([]int, n+1)
	// 3. dp数组如何初始化
	dp[0] = 1

	// 4. 确定遍历顺序
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// 2. 确定递推公式/状态转移公式
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	//// 5. debug: 打印dp数组
	//fmt.Println(dp)
	return dp[n]
}
