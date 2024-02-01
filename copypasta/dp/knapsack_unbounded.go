package dp

// 完全背包
// 更快的做法 https://www.zhihu.com/question/26547156/answer/1181239468
// https://github.com/hqztrue/shared_materials/blob/master/codeforces/101064%20L.%20The%20Knapsack%20problem%20156ms_short.cpp
// https://www.luogu.com.cn/problem/P1616
// 至少 https://www.luogu.com.cn/problem/P2918
// 恰好装满 LC322 https://leetcode-cn.com/problems/coin-change/
// EXTRA: 恰好装满+打印方案 LC1449 https://leetcode-cn.com/problems/form-largest-integer-with-digits-that-add-up-to-target/
// 【脑洞】求极限：lim_{maxW->∞} dp[maxW]/maxW
func unboundedKnapsack(values, weights []int, maxW int) int {
	dp := make([]int, maxW+1) // int64  fill
	//dp[0] = 0
	for i, cost := range values {
		w := weights[i]
		for j := w; j <= maxW; j++ {
			dp[j] = max(dp[j], dp[j-w]+cost)
		}
	}
	return dp[maxW]
}

func UnboundedKnapsack2(values, weights []int, maxW int) int {
	n := len(values)
	dp := make([][]int, n+1) // int64  fill
	for i, _ := range dp {
		dp[i] = make([]int, maxW+1)
	}
	//dp[0] = 0
	for i, cost := range values {
		w := weights[i]
		for j := 0; j <= maxW; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= w {
				dp[i+1][j] = max(dp[i][j], dp[i-1][j-w]+cost)
			}
		}
	}
	return dp[n][maxW]
}
