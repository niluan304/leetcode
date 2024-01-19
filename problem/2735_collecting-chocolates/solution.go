package main

import (
	"cmp"
	"math"
	"slices"
)

// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 10^9
// 1 <= x <= 10^9
//
// ## 提示 1
//
// 枚举操作次数，从操作 $0$ 次到操作 $n-1$ 次。
//
// ## 提示 2
//
// 如果不操作，第 $i$ 个巧克力必须花费 $\textit{nums}[i]$ 收集，总花费为所有 $\textit{nums}[i]$ 之和。例如示例 2 不操作是最优的。
//
// 如果只操作一次，第 $i$ 个巧克力可以在操作前购买，也可以在操作后购买，取最小值，即 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n])$。
//
// 如果操作两次，购买第 $i$ 个巧克力的花费为 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n], \textit{nums}[(i+2) \bmod n])$。例如示例 1，我们可以操作两次，这样每块巧克力都只需要 $1$ 的花费，总成本为 $2x+1+1+1=13$。
//
// 依此类推。
//
// ## 提示 3
//
// 如果暴力枚举操作次数，再枚举每个巧克力，再计算购买这个巧克力的最小花费，总的时间复杂度是 $\mathcal{O}(n^3)$。
//
// 一个初步的优化是，用 $\mathcal{O}(n^2)$ 的时间预处理所有子数组的最小值，保存到一个二维数组中。这样做需要 $\mathcal{O}(n^2)$ 的时间和空间。
//
// 但其实不需要预处理，还有更简单的做法：
//
// 1. 用一个长为 $n$ 的数组 $s$ 统计不同操作次数下的总成本。
// 2. 写一个二重循环，枚举子数组的左端点 $i$ 和右端点 $j$。
// 3. 在枚举右端点的同时，维护从 $\textit{nums}[i]$ 到 $\textit{nums}[j]$ 的最小值 $\textit{mn}$。
// 4. 把 $\textit{mn}$ 加到 $s[j-i]$ 中，这是因为长为 $j-i+1$ 的子数组恰好对应着操作 $j-i$ 次时要计算的子数组。
// 5. 最后输出 $\min(s)$。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(n)$。
func minCost(nums []int, x int) int64 {
	n := len(nums)
	s := make([]int64, n) // s[k] 统计操作 k 次的总成本
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形的）
			mn = min(mn, nums[j%n]) // 维护从 nums[i] 到 nums[j] 的最小值
			s[j-i] += int64(mn)     // 累加操作 j-i 次的花费
		}
	}
	return slices.Min(s)
}

// - 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(n^2)$。
func minCost2(nums []int, x int) int64 {
	n := len(nums)
	var costs = make([][]int, n)

	costs[0] = slices.Clone(nums) // 初始化，第 0 次旋转时，各巧克力的最低花费就是初始花费
	for i := 1; i < n; i++ {
		costs[i] = slices.Clone(costs[i-1])
		for j := 0; j < n; j++ {
			curCost := nums[(i+j)%n]
			costs[i][j] = min(costs[i][j], curCost) // 旋转后，巧克力的花费是否会下降？
		}
	}

	ans := math.MaxInt
	for i, cost := range costs {
		ans = min(ans, Sum(cost)+i*x)
	}
	return int64(ans)
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

// minCost2 的空间优化版本
//
// - 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(n)$。
func minCost3(nums []int, x int) int64 {
	n := len(nums)
	costs := slices.Clone(nums) // 初始化，第 0 次旋转时，各巧克力的最低花费就是初始花费

	ans := math.MaxInt
	for i := 0; i < n; i++ {
		totalCost := 0 //第 i 次旋转时，各巧克力的最低花费之和
		for j := 0; j < n; j++ {
			curCost := nums[(i+j)%n]
			costs[j] = min(costs[j], curCost) // 旋转后，巧克力的花费是否会下降？
			totalCost += costs[j]             // 累加第 j 块巧克力的最低花费
		}

		ans = min(ans, totalCost+i*x)
	}
	return int64(ans)
}

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^3)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// Deprecated: 超时
func minCost4(nums []int, x int) int64 {
	var ans = math.MaxInt
	var n = len(nums)

	// round；表示移动次数
	// 每种类型的巧克力的价格花费取所有可能位置中最小的
	for round := 0; round < n; round++ {
		sum := round * x

		// 第 i 种巧克力
		for i := 0; i < n; i++ {
			//找该种巧克力可能的最小购买值
			minCost := nums[i]

			for j := 0; j <= round; j++ {
				//进行 round 次移动时对应第 k = (i+j) mod n 位的购买价格
				k := (i + j) % n
				minCost = min(minCost, nums[k])
			}

			sum += minCost
		}

		//找最小解
		ans = min(sum, ans)
	}

	return int64(ans)
}
