package main

import "fmt"

// 1. 确定dp(dp table)数组以及下标的含义
// 2. 确定递推公式/状态转移公式
// 3. dp数组如何初始化
// 4. 确定遍历顺序
// 5. debug: 打印dp数组

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 1. 确定dp(dp table)数组以及下标的含义
	dp := make([][]int, m)
	// 3. dp数组如何初始化
	for i := range dp {
		// 跳过障碍物
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			continue
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		// 跳过障碍物
		if obstacleGrid[0][j] == 1 {
			continue
		}
		dp[0][j] = 1
	}

	// 4. 确定遍历顺序
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 2. 确定递推公式/状态转移公式
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// 5. debug: 打印dp数组
	//fmt.Println("obstacleGrid", obstacleGrid)
	//for _, rows := range dp {
	//	fmt.Println(rows)
	//}

	return dp[m-1][n-1]
}

// 1. 确定dp(dp table)数组以及下标的含义
// 2. 确定递推公式/状态转移公式
// 3. dp数组如何初始化
// 4. 确定遍历顺序
// 5. debug: 打印dp数组

// dp 动态规划
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 1. 确定dp(dp table)数组以及下标的含义
	// dp[i][j] 代表到达 i行j列的格子共有N条不同的路径
	//
	// 另外，第0行 dp[0][j] 和 第0列 dp[i][0]为不可达到的虚拟格子，即0条可达路径，这是为了方便计算 机器人只向右左 或 只向下走的情况。
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 3. 初始化dp数组
	// dp[1][1] = dp[0][1] + dp[1][0]
	// 需要将matrix[1][1]初始化为1, 因此将 dp[0][1],dp[1][0]的和设置为1, 就能达成目的
	dp[0][1] = 1

	// 4. 确定遍历顺序
	// 从前到后遍历矩阵
	// 先行后列 或 先列后行都可以
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 跳过障碍物
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			// 2. 递推公式:
			// 机器人移动到 [i][j] 的方法：
			// 在[i-1][j]处向右移动一格 或者 在[i][j-1]处向下移动一格
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// 5. debug: 打印 p数组
	fmt.Println("m int", m, "n int", n)
	for i, rows := range dp {
		fmt.Println("i", i, "rows", rows)
	}
	return dp[m][n]
}
