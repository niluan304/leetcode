package main

// dfs + 记忆化搜索
// - 时间复杂度：$\mathcal{O}(n^2 * 2)$。
// - 空间复杂度：$\mathcal{O}(n^2 * 2)$。
func stoneGameVII(stones []int) int {
	var n = len(stones)
	var ans = 0 // math.MaxInt32 // math.MinInt32

	var preSum = make([]int, n+1)
	for i, stone := range stones {
		preSum[i+1] = preSum[i] + stone
	}

	type Key struct {
		i, j    int
		isAlice bool
	}
	var cache = make(map[Key]int, n*n*2)

	var dfs func(i, j int, isAlice bool) int
	dfs = func(i, j int, isAlice bool) (res int) {
		key := Key{
			i: i, j: j,
			isAlice: isAlice,
		}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		if i == j {
			return 0
		}

		x := preSum[j] - preSum[i]
		y := preSum[j-1] - preSum[i-1]
		if isAlice {
			res = max(
				dfs(i+1, j, !isAlice)+x,
				dfs(i, j-1, !isAlice)+y,
			)
		} else {
			res = min(
				dfs(i+1, j, !isAlice)-x,
				dfs(i, j-1, !isAlice)-y,
			)
		}
		return res
	}

	ans = dfs(1, n, true)
	return ans
}

// dfs + 记忆化搜索
// - 时间复杂度：$\mathcal{O}(n^2 * 2)$。
// - 空间复杂度：$\mathcal{O}(n^2 * 2)$。
func stoneGameVII2(stones []int) int {
	var n = len(stones)
	var ans = 0 // math.MaxInt32 // math.MinInt32

	var preSum = make([]int, n+1)
	for i, stone := range stones {
		preSum[i+1] = preSum[i] + stone
	}
	var cache = make([][][2]*int, n+1)
	for i, _ := range cache {
		cache[i] = make([][2]*int, n+1)
	}

	var dfs func(i, j int, isAlice int) int
	dfs = func(i, j int, isAlice int) (res int) {
		ptr := &cache[i][j][isAlice]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		if i == j {
			return 0
		}

		x := preSum[j] - preSum[i]
		y := preSum[j-1] - preSum[i-1]
		if isAlice == 1 {
			res = max(
				dfs(i+1, j, 0)+x,
				dfs(i, j-1, 0)+y,
			)
		} else {
			res = min(
				dfs(i+1, j, 1)-x,
				dfs(i, j-1, 1)-y,
			)
		}
		return res
	}

	ans = dfs(1, n, 1)
	return ans
}

// ## 一、启发思考：寻找子问题
//
// 设爱丽丝的最终得分为 $A$，鲍勃的最终得分为 $B$，那么爱丽丝需要最大化 $A-B$，鲍勃需要最小化 $A-B$，或者说最大化 $B-A$。
//
// 也就是说，每个玩家都需要**最大化自己的得分减去对手的得分**。
//
// 看示例 1，$\textit{stones} = [5,3,1,4,2]$。枚举爱丽丝在第一回合移除的石子：
//
// - 移除最左边的石子 $\textit{stones}[0]=5$，那么需要解决的问题是：剩余石子为 $\textit{stones}' = [3,1,4,2]$，鲍勃（他现在是先手）的得分减去爱丽丝（她现在是后手）的得分最大是多少。
// - 移除最右边的石子 $\textit{stones}[4]=2$，那么需要解决的问题是：剩余石子为 $\textit{stones}' = [5,3,1,4]$，鲍勃（他现在是先手）的得分减去爱丽丝（她现在是后手）的得分最大是多少。
//
// 我们需要解决的子问题，都是**最大化「先手」的得分减去「后手」的得分**，都是**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。
//
// 虽然我们找到了子问题，但是，**原问题和子问题的关系**是什么？怎么利用子问题的结果，算出原问题的答案？
//
// ## 二、递归怎么写：状态定义与状态转移方程
//
// 因为要解决的问题都形如「对于 $\textit{stones}$ 中的一个连续子数组，计算先手得分减去后手得分的最大值」，所以定义 $\textit{dfs}(i,j)$ 表示剩余石子从 $\textit{stones}[i]$ 到 $\textit{stones}[j]$，先手得分减去后手得分的最大值。
//
// 例如 $\textit{stones} = [5,3,1,4,2]$。在第一回合，如果爱丽丝移除最右边的 $\textit{stones}[4]=2$，得到 $\textit{pt}_4 = 5 + 3 + 1 + 4 = 13$ 分，那么问题变成：对于 $\textit{stones}' = [5,3,1,4]$，鲍勃选哪颗石子，可以最大化「鲍勃的得分减去爱丽丝的得分」，注意这里的得分是指在 $\textit{stones}'$ 上的得分
// 。
//
// 对于 $\textit{stones}'$，设鲍勃的最终得分为 $B'$，爱丽丝的最终得分为 $A'$，那么子问题
//
// $$
// \textit{dfs}(0,3) = B'-A'
// $$
//
// 我们要计算的原问题
//
// $$
// \textit{dfs}(0,4)=A-B
// $$
//
// 由于 $A = \textit{pt}_4 + A',\ B = B'$，所以有
//
// $$
// \textit{dfs}(0,4) = A-B = (\textit{pt}_4+A') - B' = \textit{pt}_4 - (B'-A') = \textit{pt}_4 - \textit{dfs}(0,3)
// $$
//
// 即
//
// $$
// \textit{dfs}(0,4) = \textit{pt}_4 - \textit{dfs}(0,3)
// $$
//
// 这样就找到了**原问题和子问题的关系**。
//
// 一般地，如果剩余石子从 $\textit{stones}[i]$ 到 $\textit{stones}[j]$，我们枚举先手移除的石子：
//
// - 如果移除的是最左边的石子 $\textit{stones}[i]$，利用 $\textit{stones}$ 的前缀和 $s$ 可以算出，我们得到了 $s[j+1]-s[i+1]$ 分，这种情况下 $\textit{dfs}(i,j) = s[j+1]-s[i+1]-\textit{dfs}(i+1,j)$。
// - 如果移除的是最右边的石子 $\textit{stones}[j]$，我们得到了 $s[j]-s[i]$ 分，这种情况下 $\textit{dfs}(i,j) = s[j]-s[i]-\textit{dfs}(i,j-1)$。
//
// 这两种情况取最大值，即
//
// $$
// \textit{dfs}(i,j) = \max(s[j+1]-s[i+1]-\textit{dfs}(i+1,j),\ s[j]-s[i]-\textit{dfs}(i,j-1))
// $$
//
// 递归边界：$\textit{dfs}(i,i)=0$。此时只有一颗石子 $\textit{stones}[i]$，移除后没有剩余石子，无法继续游戏，得分之差为 $0$。
//
// 递归入口：$\textit{dfs}(0,n-1)$，也就是答案。
//
// > 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「枚举选（移除）哪个」。
func stoneGameVII3(stones []int) int {
	n := len(stones)
	preSum := make([]int, n+1) // 前缀和
	for i, x := range stones {
		preSum[i+1] = preSum[i] + x
	}
	cache := make([][]*int, n)
	for i := range cache {
		cache[i] = make([]*int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == j { // 递归边界
			return 0
		}
		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		x := preSum[j+1] - preSum[i+1]
		y := preSum[j] - preSum[i]
		return max(
			x-dfs(i+1, j),
			y-dfs(i, j-1),
		)
	}

	return dfs(0, n-1)
}

// 将 stoneGameVII3 1:1 翻译为递推
// #### 答疑
//
// **问**：
// 如何思考循环顺序？什么时候要正序枚举，什么时候要倒序枚举？
//
// **答**：
// 这里有一个通用的做法：盯着状态转移方程，想一想，要计算 $f[i][j]$，
// 其一，必须先把 $f[i+1][j]$ 算出来，那么只有 $i$ 从大到小枚举才能做到；
// 其二，必须先把 $f[i][j-1]$ 算出来，那么对于内层循环，只有 $j$ 从小到大枚举才能做到。
//
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n^2)$。
func stoneGameVII4(stones []int) int {
	n := len(stones)
	preSum := make([]int, n+1) // 前缀和
	for i, x := range stones {
		preSum[i+1] = preSum[i] + x
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			x := preSum[j+1] - preSum[i+1]
			y := preSum[j] - preSum[i]

			dp[i][j] = max(
				x-dp[i+1][j],
				y-dp[i][j-1],
			)
		}
	}

	return dp[0][n-1]
}
