package main

// 方法一：动态规划
// 爱丽丝获胜的概率只和下一轮开始前的得分有关，因此根据得分计算概率。
//
// 令 dp[x]  表示从得分为 x 的情况开始游戏并且获胜的概率，目标是求 dp[0] 的值。
//
// 根据规则，当分数达到或超过 k 时游戏结束，游戏结束时，如果分数不超过 n 则获胜，
// 如果分数超过 n 则失败。因此当 k≤x≤min(n,k+maxPts−1)时有 dp[x]=1，
// 当 x>min(n,k+maxPts−1) 时有 dp[x]=0。
//
// 为什么分界线是 min(n,k+maxPts−1)？首先，只有在分数不超过 n 时才算获胜；
// 其次，可以达到的最大分数为 k+maxPts−1，即在最后一次抽取数字之前的分数为 k−1，并且抽到了 maxPts。
//
// 当 0≤x<k 时，如何计算 dp[x] 的值？注意到每次在范围 [1,maxPts] 内随机抽取一个整数，
// 且每个整数被抽取到的概率相等，因此可以得到如下状态转移方程：
func leetcode1(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}
	for i := k - 1; i >= 0; i-- {
		for j := 1; j <= maxPts; j++ {
			dp[i] += dp[i+j] / float64(maxPts)
		}
	}
	return dp[0]
}

// 上述解法的时间复杂度是 O(n+k×maxPts)，会超出时间限制，因此需要优化。
func leetcode2(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}

	dp[k-1] = 1.0 * float64(min(n-k+1, maxPts)) / float64(maxPts)
	for i := k - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+maxPts+1]-dp[i+1])/float64(maxPts)
	}
	return dp[0]
}
