package main

import (
	"slices"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// dp
//
// LC2919 https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/
//
// - 时间复杂度：$\mathcal{O}(n \cdot k)$。
// - 空间复杂度：$\mathcal{O}(n \cdot k)$。
// Deprecated: 超时
func bruteForce(nums []int, k int) int {

	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= n {
			return 0
		}

		res := dfs(i+1, 1) + nums[i] // 选
		if j < k {
			res = max(res, dfs(i+1, j+1)) // 未选数不足 k，可以不选
		}
		return res
	}

	cache := MemorySearch2(&dfs)
	_ = cache
	ans = dfs(0, 1)
	return ans
}

// - 时间复杂度：$\mathcal{O}(n \cdot k)$。
// - 空间复杂度：$\mathcal{O}(n \cdot k)$。
// Deprecated: 超时
func bruteForce2(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = slices.Max(dp[max(i-k, 0):i]) + nums[i]
	}
	return dp[n-1]
}

// 单调队列优化的 dp
//
// 借助单调队列，做到 O(1) 时间内获取 bruteForce2 的 Max(dp[max(i-k, 0):i]) 的值
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maxResult(nums []int, k int) int {
	var queue []int
	n := len(nums)
	dp := make([]int, n)

	for i, num := range nums {
		for len(queue) > 0 && i-queue[0] > k { // 去除过期元素
			queue = queue[1:]
		}

		dp[i] = num // 手动初始化可以优化代码结构
		if i > 0 {  // 等价于 len(queue) > 0
			dp[i] += dp[queue[0]] // 队首就是 [i-k:i-1] 中的最大得分位置，对应着 Max(dp[max(i-k, 0):i])
		}

		m := len(queue)
		for m > 0 && dp[i] >= dp[queue[m-1]] {
			m--
		}
		queue = append(queue[:m], i) // 新元素新队，维护单调性
	}

	return dp[n-1]
}

// 单调队列优化的 dp
//
// maxResult 的手动初始化版本
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maxResult2(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	queue := []int{0}

	for i := 1; i < n; i++ {
		// len(queue) 必定 > 0
		for i-queue[0] > k { // 去除过期元素
			queue = queue[1:]
		}
		dp[i] = dp[queue[0]] + nums[i] // 队首就是 [i-k:i-1] 中的最大得分位置，对应着 Max(dp[max(i-k, 0):i])

		m := len(queue)
		for m > 0 && dp[i] >= dp[queue[m-1]] {
			m--
		}
		queue = append(queue[:m], i) // 新元素新队，维护单调性
	}
	return dp[n-1]
}
