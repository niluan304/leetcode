package dp

import (
	"math"
)

// 0-1背包常见变形：
// - 至多装capacity，求方案数/最大价值和
// - 恰好装capacity，求方案数/最大/最小价值和
// - 至少装capacity，求方案数/最小价值和

// ZeroOneKnapsack
// 0-1 背包 (n 个物品，背包容量为 maxW，每个物品只能选 1 次)
// 状态：从前 i 个物品中选择若干个，当容量限制为 j 时能获得的最大价值和  i∈[0,n-1], j∈[0,maxW]
// 初始值：f(0,j)=0  j∈[0,maxW]
// 除开初始状态，每个状态有两个来源，决策为 max：
// - 不选第 i 个物品：f(i-1,j) -> f(i,j)
// - 　选第 i 个物品：f(i-1,j-wi)+vi -> f(i,j)  j ≥ wi
//
// 最优解为 f(n-1,maxW)
//
// https://oi-wiki.org/dp/knapsack/
func ZeroOneKnapsack(values, weights []int, maxW int) int {
	// 背包的最大价值和
	n := len(weights)
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i := 0; i < n; i++ {
		v := values[i]
		for j := 0; j <= maxW; j++ {
			dp[i+1][j] = dp[i][j] // 不选当前物品
			if w := j - weights[i]; w >= 0 {
				dp[i+1][j] = max(dp[i+1][j], v+dp[i][w]) // 选择当前物品
			}
		}
	}

	return dp[n][maxW]
}

// ZeroOneKnapsackExactlyFull
// 01背包变形：恰好装满
// https://leetcode.cn/contest/sf-tech/problems/cINqyA/
// LC 2915 https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/
// LC 416 https://leetcode.cn/problems/partition-equal-subset-sum/
func ZeroOneKnapsackExactlyFull(values, weights []int, maxW int) int {
	// 背包的最大价值和
	n := len(values)
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, maxW+1)
		for j, _ := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}

	dp[0][0] = 0
	for i := 0; i < n; i++ {
		v := values[i]

		for j := 0; j <= maxW; j++ {
			dp[i+1][j] = dp[i][j] // 不选当前物品
			if k := j - weights[i]; k >= 0 {
				dp[i+1][j] = max(dp[i+1][j], v+dp[i][k]) // 选择当前物品
			}
		}
	}

	for i := n; i >= 0; i-- {
		if dp[i][maxW] >= 0 { // 能恰好装满 maxW，此时背包物品价值和的最大值是 dp[i][maxW]
			return dp[i][maxW]
		}
	}
	return math.MinInt32
}

// ZeroOneKnapsackAtLeastFillUp
// 0-1 背包 EXTRA: 至少装入重量和为 maxW 的物品，求价值和的最小值 https://www.luogu.com.cn/problem/P4377
// f[0] 表示至少为 0 的情况，也表示没有任何约束的情况
// 比如选第 i 个物品后容量 <=0 了，那就表示前面的 i-1 个物品可以不受约束地随意选或不选了
// - 转换
// https://codeforces.com/problemset/problem/19/B
// LC2742 https://leetcode.cn/problems/painting-the-walls/
// 二维费用的情况+价值最小 https://ac.nowcoder.com/acm/contest/6218/C
func ZeroOneKnapsackAtLeastFillUp(values, weights []int, maxW int) int {
	n := len(weights)
	dp := make([][]int, n+1) // int64
	for i := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i, _ := range dp[0] {
		dp[0][i] = math.MaxInt32
	}
	dp[0][0] = 0

	for i, value := range values {
		w := weights[i]
		for j := 0; j <= maxW; j++ { // 0-1背包：压缩 dp 数组时，需要倒序遍历，why?
			k := max(j-w, 0) // max(j-w, 0) 蕴含了「至少」
			dp[i+1][j] = min(dp[i][j], dp[i][k]+value)
		}
	}

	//// 另一种写法
	//for i, value := range values {
	//	w := weights[i]
	//	for j := maxW; j >= 0; j-- {
	//		k := min(j+w, maxW)
	//		dp[i+1][k] = min(dp[i][k], dp[i][j]+value)
	//	}
	//}

	return dp[n][maxW]
}

// ZeroOneWaysToSum
// 0-1 背包 EXTRA: 从序列 a 中选若干个数，使其总和为 sum 的方案数
// 常见题目是算严格分拆（必须用不同数字）
// LC2787 https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/
// - https://oeis.org/A000009
// NOTE: 1,1,1,...1(32个1),s-32,s-31,...,s 可以让方案数恰好为 2^32
// 二维+上限+下限 LC879 https://leetcode.cn/problems/profitable-schemes/
// https://atcoder.jp/contests/arc060/tasks/arc060_a
// https://codeforces.com/problemset/problem/1673/C
//
// - 转换
// https://atcoder.jp/contests/abc169/tasks/abc169_f
// https://codeforces.com/problemset/problem/478/D
// LC494 https://leetcode.cn/problems/target-sum/
// LC1434 https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other/
// 由于顺序不同也算方案，所以这题需要正序递推 LC377 https://leetcode.cn/problems/combination-sum-iv/
func ZeroOneWaysToSum(nums []int, sum int) int {
	n := len(nums)
	dp := make([][]int, n+1) // int64
	for i, _ := range dp {
		dp[i] = make([]int, sum+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for s := 0; s <= sum; s++ {
			dp[i+1][s] = dp[i][s]
			if s >= num {
				dp[i+1][s] += dp[i][s-num] // % mod
			}
		}
	}
	return dp[n][sum]
}

// 完全背包常见变形：
// - 至多装capacity，求方案数/最大价值和
// - 恰好装capacity，求方案数/最大/最小价值和
// - 至少装capacity，求方案数/最小价值和

// UnboundedKnapsack
// 完全背包
// 更快的做法 https://www.zhihu.com/question/26547156/answer/1181239468
// https://github.com/hqztrue/shared_materials/blob/master/codeforces/101064%20L.%20The%20Knapsack%20problem%20156ms_short.cpp
// https://www.luogu.com.cn/problem/P1616
// 至少 https://www.luogu.com.cn/problem/P2918
// 恰好装满 LC322 https://leetcode.cn/problems/coin-change/
// EXTRA: 恰好装满+打印方案 LC1449 https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/
// 【脑洞】求极限：lim_{maxW->∞} dp[maxW]/maxW
func UnboundedKnapsack(values, weights []int, maxW int) int {
	n := len(values)
	dp := make([][]int, n+1) // int64  fill
	for i, _ := range dp {
		dp[i] = make([]int, maxW+1)
	}
	//dp[0] = 0
	for i, value := range values {
		w := weights[i]
		for j := 0; j <= maxW; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= w {
				dp[i+1][j] = max(dp[i+1][j], dp[i+1][j-w]+value)
			}
		}
	}
	return dp[n][maxW]
}

// UnboundedKnapsackExactlyFull
// 完全背包变形：恰好装满时的 最大价值和 / 最小价值和 /方案数
// 最小价值和：LC322 https://leetcode.cn/problems/coin-change/
// 方案数：   LC518 https://leetcode.cn/problems/coin-change-ii/
func UnboundedKnapsackExactlyFull(values, weights []int, maxW int) int {
	n := len(values)
	dp := make([][]int, n+1) // int64  fill
	for i, _ := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i, _ := range dp[0] {
		dp[0][i] = math.MinInt32
	}

	dp[0][0] = 0
	for i, value := range values {
		w := weights[i]
		for j := 0; j <= maxW; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= w {
				dp[i+1][j] = max(dp[i+1][j], dp[i+1][j-w]+value)
			}
		}
	}
	ans := dp[n][maxW]
	if ans < 0 { // 无法恰好装满背包
		return -1
	}
	return ans
}

// UnboundedKnapsackAtLeastFillUp
// 至少装入重量和为 maxW 的物品，求价值和的最小值
// https://www.luogu.com.cn/problem/P2918
func UnboundedKnapsackAtLeastFillUp(values, weights []int, maxW int) int {
	n := len(weights)
	dp := make([][]int, n+1) // int64
	for i := range dp {
		dp[i] = make([]int, maxW+1)
	}
	for i, _ := range dp[0] {
		dp[0][i] = math.MaxInt32
	}
	dp[0][0] = 0

	for i, value := range values {
		w := weights[i]
		for j := 0; j <= maxW; j++ { // 完全背包：压缩 dp 数组时，需要正序遍历，why?
			k := max(j-w, 0) // max(j-w, 0) 蕴含了「至少」
			dp[i+1][j] = min(dp[i][j], dp[i+1][k]+value)
		}
	}

	return dp[n][maxW]
}

// UnboundedKnapsackWaysToSum
// 完全背包，恰好装满背包时的方案数
// LC518 https://leetcode.cn/problems/coin-change-ii/
func UnboundedKnapsackWaysToSum(nums []int, sum int) int {
	n := len(nums)
	dp := make([][]int, n+1) // int64  fill
	for i, _ := range dp {
		dp[i] = make([]int, sum+1)
	}

	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= sum; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i+1][j-num]
			}
		}
	}
	ans := dp[n][sum]
	return ans
}