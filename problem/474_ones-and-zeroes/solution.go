package main

import "strings"

// dp, 动态规划
// 时间复杂度：O(lmn)
// 空间复杂度：O(lmn)
func findMaxForm(strs []string, m int, n int) int {
	var l = len(strs)
	var dp = make([][][]int, l+1)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i, str := range strs {
		// 计算当前物品的价值 [w0,w1]
		w0 := strings.Count(str, "0")
		w1 := len(str) - w0

		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				// 递推公式：选与不选 价值为[w0,w1]的当前物品
				if j < w0 || k < w1 {
					dp[i+1][j][k] = dp[i][j][k]
				} else {
					dp[i+1][j][k] = max(dp[i][j][k], dp[i][j-w0][k-w1]+1)
				}
			}
		}
	}

	//// debug: 打印dp数组
	//for i := range dp {
	//	for j := range dp[i] {
	//		fmt.Println(dp[i][j])
	//	}
	//	fmt.Println()
	//}
	return dp[l][m][n]
}

// dp 动态规划, 01 背包, 复用二维数组, 以降低空间复杂度
// 时间复杂度：O(lmn)
// 空间复杂度：O(mn)
func findMaxForm2(strs []string, m int, n int) int {
	var dp = make([][]int, m+1)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		// 计算当前物品的价值 [w0,w1]
		w0 := strings.Count(str, "0")
		w1 := len(str) - w0
		for j := m; j >= w0; j-- {
			for k := n; k >= w1; k-- {
				// 递推公式：选与不选 价值为[w0,w1]的当前物品
				dp[j][k] = max(dp[j][k], dp[j-w0][k-w1]+1)
			}
		}

		//// debug: 打印dp数组
		//for i := range dp {
		//	fmt.Println(dp[i])
		//}
		//fmt.Println()
	}

	return dp[m][n]
}
