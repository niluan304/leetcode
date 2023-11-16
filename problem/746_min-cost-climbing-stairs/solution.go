package main

func minCostClimbingStairs(cost []int) int {
	// 1. 确定dp数组以及下标的含义
	// dp[i]表示第i个台阶锁需要支付的最小费用
	var (
		n  = len(cost)
		dp = make([]int, n+1)
	)

	// 3. 初始化dp数组
	// 可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯，
	// 表示 dp[0] 和 dp[1] 都为0
	dp[0] = 0
	dp[1] = 0
	// 4. 遍历顺序
	// 从小到大遍历 cost
	for i := 2; i <= n; i++ {
		// 2. 确定递推公式
		// 第i个台阶只能从 i-1 或 i-2 中爬上来，比较两者的花费，即可得到第i个台阶的最小费用
		dp[i] = _min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	// 5. debug: 打印dp数组
	//fmt.Println(dp)
	return dp[n]
}

func _min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func minCostClimbingStairs2(cost []int) int {
	// 1. 确定dp数组以及下标的含义
	// dp[i]表示第i个台阶锁需要支付的最小费用
	//var (
	//	n  = len(cost)
	//	dp = make([]int, n+1)
	//)

	var (
		// 3. 初始化dp数组
		// 可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯，
		// 表示 dp[0] 和 dp[1] 都为0
		dp0 = 0
		dp1 = 0
		dpi = 0
	)
	// 4. 遍历顺序
	// 从小到大遍历 cost
	for i := 2; i <= len(cost); i++ {
		// 2. 确定递推公式
		// 第i个台阶只能从 i-1 或 i-2 中爬上来，比较两者的花费，即可得到第i个台阶的最小费用
		dpi = _min(dp1+cost[i-1], dp0+cost[i-2])
		dp1, dp0 = dpi, dp1
	}

	// 5. debug: 打印dp数组
	//fmt.Println(dp)
	return dpi
}
