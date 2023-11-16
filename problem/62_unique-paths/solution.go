package main

// dp 动态规划
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func uniquePaths(m int, n int) int {
	// 1. 确定dp(dp table)数组以及下标的含义
	// 2. 确定递推公式/状态转移公式
	// 3. dp数组如何初始化
	// 4. 确定遍历顺序
	// 5. debug: 打印dp数组

	// 1. 确定dp(dp table)数组以及下标的含义
	// matrix[i][j] 代表到达 i行j列的格子共有N条不同的路径
	//
	// 另外，第0行 matrix[0][j] 和 第0列 matrix[i][0]为不可达到的虚拟格子，即0条可达路径，这是为了方便计算 机器人只向右左 或 只向下走的情况。
	var matrix = make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	// 3. 初始化dp数组
	// matrix[1][1] = matrix[0][1] + matrix[1][0]
	// 需要将matrix[1][1]初始化为1, 因此将 matrix[0][1],matrix[1][0]的和设置为1, 就能达成目的
	matrix[0][1] = 1

	// 4. 确定遍历顺序
	// 从前到后遍历矩阵
	// 先行后列 或 先列后行都可以
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 2. 递推公式:
			// 机器人移动到 [i][j] 的方法：
			// 在[i-1][j]处向右移动一格 或者 在[i][j-1]处向下移动一格
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	// 5. debug: 打印 dp数组
	//fmt.Println("m int", m, "n int", n)
	//for i, rows := range matrix {
	//fmt.Println("i", i, "rows", rows)
	//}

	return matrix[m][n]
}

// 递归
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func uniquePaths2(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return uniquePaths2(m-1, n) + uniquePaths2(m, n-1)
}

// dp 动态规划
// 时间复杂度：O(mn)
// 空间复杂度：O(min(m,n))
// var matrix = make([][]int, m+1) 的优化版本
// 空间复杂度 从O(mn)降低至 O(min(m,n))
func uniquePaths3(m int, n int) int {
	// m行n列　与 n行m列 等价, 使用 _min(m,n)可节约空间
	if n > m {
		m, n = n, m
	}
	// 1. 确定dp(dp table)数组以及下标的含义
	// 从左向右遍历
	var rows = make([]int, n)

	// 3. 初始化dp数组
	// 首列的格子只有一种可能
	rows[0] = 1

	// 4. 确定遍历顺序
	// 逐列遍历
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			// 2. 递推公式:
			// 机器人移动到 [i][j] 的方法：
			// 在[i-1][j]处向右移动一格 或者 在[i][j-1]处向下移动一格
			// 更新前 rows[j] 表示左边的格子, [j-1] 表示上方的格子,
			// 两者相加，即更新后表示机器人完成移动
			rows[j] = rows[j] + rows[j-1]
		}
		//  5. debug: 打印 dp数组
		//fmt.Println("i", i, "rows", rows)
	}

	return rows[n-1]
}
