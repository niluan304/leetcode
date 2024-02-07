package main

func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}

	var ans = 0
	// 将 n分为 k 段, 题目要求 k >= 2
	for k := 2; k <= n/2; k++ {
		// 和固定，k也固定时，每段值越接近，乘积也就越大
		product := 1      // 本次切割的乘积
		item := n / k     // 计算每段的值
		num := n - k*item // 计算余数
		for i := 0; i < num; i++ {
			product *= item + 1
		}
		for i := num; i < k; i++ {
			product *= item
		}

		ans = max(ans, product)
	}
	return ans
}

// 1. 确定dp(dp table)数组以及下标的含义
// 2. 确定递推公式/状态转移公式
// 3. dp数组如何初始化
// 4. 确定遍历顺序
// 5. debug: 打印dp数组

func integerBreak2(n int) int {
	// 1. 确定dp(dp table)数组以及下标的含义
	// dp[i] 第i位的最大乘积
	var dp = make([]int, n+1)

	// 3. dp数组如何初始化
	// dp[0], dp[1]无意义,
	// dp[2] = 1x1 = 1,
	// dp[3] = 1x2 = 2
	dp[2] = 1

	// 4. 确定遍历顺序
	// 从小到大遍历
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// 2. 确定递推公式/状态转移公式
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	//// 5. debug: 打印dp数组
	//fmt.Println(dp)
	return dp[n]
}
