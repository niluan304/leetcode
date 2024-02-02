package main

import (
	. "github.com/niluan304/leetcode/copypasta/dp"
)

// 完全背包
//
// - 时间复杂度：$\mathcal{O}(n \codt m)$。
// - 空间复杂度：$\mathcal{O}(n \codt m)$。
// 和爬楼梯的区别？
// 需要 amount 阶你才能到达楼顶。你有 coins[:] 方式向上爬
// 你从 0 开始出发，有多少种方式爬到楼顶？
//
// 「排列」 与 「组合」的区别
//
//   - 爬楼梯是排列
//     对于 [1,2,1] 与 [1,1,2] 是不同的方案
//
//   - 组硬币是组合
//     对于 [1,2,1] 与 [1,1,2] 是相同的方案
func change(amount int, coins []int) int {
	var dfs func(i int, amount int) int
	dfs = func(i int, j int) int {
		if i == 0 {
			if j == 0 {
				return 1
			}
			return 0
		}

		res := dfs(i-1, j) // 不选 coins[i]
		if k := j - coins[i-1]; k >= 0 {
			res += dfs(i, k) // 选择 coins[i]
		}
		return res
	}

	MemorySearch2(&dfs)
	ans := dfs(len(coins), amount)
	return ans
}

// 完全背包
//
// - 时间复杂度：$\mathcal{O}(n \codt m)$。
// - 空间复杂度：$\mathcal{O}(n \codt m)$。
func change2(amount int, coins []int) int {
	return UnboundedKnapsackWaysToSum(coins, amount)
}
