package main

import (
	"math"
)

// dp 动态规划, 完全背包
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func numSquares(n int) int {
	var c = int(math.Sqrt(float64(n))) // 背包里最大值明显为 n开平方后向下取证的整数, 即 math.Ceil(math.Sqrt(n))
	var dp = make([][]int, c+1)        // dp[i][j] 表示 恰好填满容量为j的背包时，从[0,i]种物品中，凑出方案的最小值
	for i := 0; i <= c; i++ {
		dp[i] = make([]int, n+1)
		// 初始化 dp数组
		for j := 1; j < n+1; j++ {
			dp[i][j] = math.MaxInt32
		}
		// 平方数只需要1个
		dp[i][i*i] = 1
	}

	for i := 1; i <= c; i++ {
		w := i * i // 第 i 件物品，重量为 i, 价值为 i*i
		for j := 1; j <= n; j++ {
			// 递推公式, 在背包的总空间为 j 情况，选与不选 第 i 件物品
			dp[i][j] = _min(dp[i][j], dp[i-1][j]) // 不选 w
			if j > w {
				dp[i][j] = _min(dp[i][j], dp[i][j-w]+1) // 选 w
			}
		}
	}

	//// debug: 打印 dp数组
	//for i := range dp {
	//	fmt.Println("n", n, "i", i, dp[i])
	//}

	return dp[c][n]
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// dp 动态规划, 完全背包, 复用一维数组, 以降低空间复杂度
// 时间复杂度：O(m*n)
// 空间复杂度：O(n)
func numSquares2(n int) int {
	var c = int(math.Sqrt(float64(n)))
	var dp = make([]int, n+1)

	// 初始化 dp数组
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32 // 用于 Min判断
	}
	for i := 1; i <= c; i++ {
		w := i * i // 第 i 件物品，重量为 i, 价值为 i*i
		dp[w] = 1
		for j := w + 1; j <= n; j++ {
			// 递推公式, 在背包的总空间为 i 情况，选与不选 第 i 件物品
			dp[j] = _min(dp[j], dp[j-w]+1) // 选 w
		}

		//// debug: 打印 dp数组
		//fmt.Println("n", n, "i", i, dp)
	}

	return dp[n]
}
