package main

import (
	"math"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// ### 视频讲解
//
// 见[【周赛 350】](https://www.bilibili.com/video/BV1Hj411D7Tr/)第四题，欢迎点赞投币！
//
// ## 方法一：选或不选
//
// ### 前置知识：动态规划入门
//
// 详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)
//
// ### 思路
//
// 用「选或不选」的思路来思考：
//
// - 如果付费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，且付费时间和为 $\textit{time}[n-1]$，免费时间和 $0$ 的最少开销（开销指 $\textit{cost}$ 的和）。
// - 如果免费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，且付费时间和为 $0$，免费时间和为 $1$ 的最少开销。
//
// 因此，定义 $\textit{dfs}(i,j,k)$ 表示刷前 $i$ 堵墙，且付费时间和为 $j$，免费时间和为 $k$ 的最少开销。递归到终点时，如果 $j\ge k$，说明这种方案是合法的，否则不合法。
//
// 但是，这样定义的话，状态个数就太多了，需要优化。由于最后是比较的 $j$ 和 $k$ 的「相对大小」，那么不妨把 $j$ 重新定义为【付费时间和】减去【免费时间和】，这样递归到终点时，如果 $j\ge 0$，说明这种方案是合法的，否则不合法。这样一来，状态个数就大大减少了。
//
// 分类讨论：
//
// - 如果付费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i]$。
// - 如果免费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j-1)$。
//
// 两种情况取最小值：
//
// $$
// \textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i], \textit{dfs}(i-1,j-1))
// $$
//
// 递归边界：如果 $j>i$，那么剩余的墙都可以免费刷，即 $\textit{dfs}(i,j) = 0$，否则 $\textit{dfs}(-1,j) = \infty$。
//
// 递归入口：$\textit{dfs}(n-1,0)$。
func paintWalls(cost, time []int) int {
	n := len(cost)
	cache := make([][]*int, n)
	for i := range cache {
		cache[i] = make([]*int, n*2+1)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j-n > i { // 注意 j 加上了偏移量 n
			return 0 // 剩余的墙都可以免费刷
		}
		if i < 0 {
			return math.MaxInt32 // 防止加法溢出
		}
		p := &cache[i][j]
		if *p != nil {
			return **p
		}
		defer func() { *p = &res }()

		return min(
			dfs(i-1, j+time[i])+cost[i], // 付费刷油漆
			dfs(i-1, j-1),               // 免费刷油漆
		)
	}
	return dfs(n-1, n) // 加上偏移量 n，防止负数
}

func paintWalls2(cost []int, time []int) int {
	n := len(cost)

	var cache = make([][]*int, n)
	for i, _ := range cache {
		cache[i] = make([]*int, n+1)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if j <= 0 {
			// 剩余需要的体积为 0，没有约束，因此最小的价值和为 0，不选物品
			return 0
		}
		if i < 0 {
			return math.MaxInt32
		}

		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		return min(
			dfs(i-1, j),                     // 不选：至少装 j 时, 前 i-1 最小价值和
			dfs(i-1, j-(time[i]+1))+cost[i], //   选：容量变小，问题变为：至少装 j-(time[i]+1) 时，前 i-1 物品的最小价值和
		)
	}

	ans := dfs(n-1, n)
	return ans
}

func paintWalls3(cost []int, time []int) int {
	n := len(cost)
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, _ := range dp[0] {
		dp[0][i] = math.MaxInt32 // 体积需求大于零，却没有物品可选，是无法达成的，此时最小价值和设置为无穷大
	}
	dp[0][0] = 0

	// 把 time[i] + 1 看成物品体积
	// cost[i] 看成物品价值
	// 从 n 个物品中选择体积和「至少」为 n，价值和最小是多少？
	maxW := n
	for i, value := range cost {
		w := time[i] + 1
		for j := 0; j <= maxW; j++ {
			k := max(0, j-w) // 剩余需要的体积如果小于等于 0，意味着没有约束，因此最小的价值和为 0，不选物品

			dp[i+1][j] = min(
				dp[i][j],       // 不选：至少装 j 时, 前 i-1 最小价值和
				dp[i][k]+value, //   选：容量变小，问题变为：至少装 j-(time[i]+1) 时，前 i-1 物品的最小价值和
			)
		}
	}
	return dp[n][maxW]
}

// ## 方法二：转换成 0-1 背包
//
// ### 前置知识：0-1 背包
//
// 详见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)
//
// ### 思路
//
// 看着方法一的状态转移方程，你可能会觉得：怎么感觉很像 0-1 背包呢？没错，这题就是 0-1 背包！
//
// 根据题意，付费刷墙个数 $+$ 免费刷墙个数 $=n$。
//
// 同时，付费刷墙时间之和必须 $\ge$ 免费刷墙个数。
//
// 结合这两个式子，得到：付费刷墙时间之和 $\ge n\ -$ 付费刷墙个数。
//
// 移项，得到：「付费刷墙时间$+1$」之和 $\ge n$。（把个数拆分成 $1+1+1+\cdots$。）
//
// 把 $\textit{time}[i]+1$ 看成物品体积，$\textit{cost}[i]$ 看成物品价值，问题变成：
//
// 从 $n$ 个物品中选择体积和**至少**为 $n$ 的物品，价值和最小是多少？
//
// 这是 0-1 背包的一种「至少装满」的变形。我们可以定义 $\textit{dfs}(i,j)$ 表示考虑前 $i$ 个物品，**剩余**还需要凑出 $j$ 的体积，此时的最小价值和。
//
// 依然是用「选或不选」来思考，可以得到类似的状态转移方程：
//
// $$
// \textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j-\textit{time}[i]-1)+\textit{cost}[i], \textit{dfs}(i-1,j))
// $$
//
// 递归边界：如果 $j\le 0$，那么不需要再选任何物品了，返回 $0$；如果 $i<0$，返回无穷大。
//
// 递归入口：$\textit{dfs}(n-1,n)$，表示体积和至少为 $n$，这正是我们要计算的。
func paintWalls4(cost []int, time []int) int {
	for i, _ := range time {
		time[i]++
	}
	return ZeroOneKnapsackAtLeastFillUp(cost, time, len(cost))
}
