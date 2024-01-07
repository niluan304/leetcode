package main

import (
	"math"
	"sort"
)

// ## 题意（版本 A）
//
// 说，有个小偷公司，给小偷定的 KPI 是偷至少 $k$ 间房子，要求偷的房子不能相邻。
//
// 张三作为其中的一个小偷，他不想偷太多，否则一不小心就「数额巨大」，这可太刑了。所以张三计划，在他偷过的房子中，偷走的最大金额要尽量地小。
//
// 这个最小值是多少呢？
//
// ## 题意（版本 B）
//
// 给定数组 $\textit{nums}$，从中选择一个长度至少为 $k$ 的子序列 $A$，要求 $A$ 中没有任何元素在 $\textit{nums}$ 中是相邻的。
//
// 最小化 $\max(A)$。
//
// ## 方法一：二分+DP
//
// 有关二分的三种写法，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。本文采用**开区间**写法。
//
// 看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
// - 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
func minCapability(nums []int, k int) int {
	var n = len(nums)
	var dp = make([]int, n+2)
	var rob = func(ability int) bool {
		for i, num := range nums {
			dp[i+2] = dp[i+1] // dp[0], dp[1] 始终为 0, 会把后面的值覆盖，就不用每次都 clear dp数组了。
			if num <= ability {
				dp[i+2] = max(dp[i+2], dp[i]+1)
			}
		}
		return dp[n+1] >= k
	}
	return sort.Search(1e9, rob)
}

// dp, 动态规划, 本质为爬楼梯
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
// Deprecated: 超时
func minCapability2(nums []int, k int) int {
	var rob = func(capability int) bool {
		var dp0, dp1 = 0, 0
		for _, num := range nums {
			dp2 := dp1
			if capability >= num {
				dp2 = max(dp2, dp0+1)
			}
			dp1, dp0 = dp2, dp1
		}
		return dp1 >= k
	}

	var ans = math.MaxInt
	for _, num := range nums {
		if num > ans {
			continue
		}
		if rob(num) {
			ans = min(ans, num)
		}
	}
	return ans
}
