package main

import (
	"slices"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

const Len = 3 // 题目要求：如果数组中任何长度 大于或等于 3 的子数组，因此 Len = 3

// dfs + 记忆化搜索
//
// - 时间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minIncrementOperations(nums []int, k int) int64 {
	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i < 0 {
			return 0
		}

		// 如果把当前数字调整为不小于 k 的数需要的操作数量为 max(0, k - num)，而此前需要的操作数量为 min(dp1, dp2, dp3)
		// 原来距离为 1 变为距离为 2
		// 原来距离为 2 变为距离为 3
		return max(0, k-nums[i]) + min(dfs(i-1), dfs(i-2), dfs(i-3))
	}

	cache := MemorySearch(&dfs)
	_ = cache

	n := len(nums)
	ans := min(dfs(n-1), dfs(n-2), dfs(n-3))

	return int64(ans)
}

// DP
//
// 将 minIncrementOperations 1:1 翻译为递推
//
// - 时间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minIncrementOperations2(nums []int, k int) int64 {
	var n = len(nums)
	x := Len
	var dp = make([]int, n+x)
	for i, num := range nums {
		// 如果把当前数字调整为不小于 k 的数需要的操作数量为 max(0, k - num)，而此前需要的操作数量为 min(dp1, dp2, dp3)
		// 原来距离为 1 变为距离为 2
		// 原来距离为 2 变为距离为 3

		dp[i+x] = max(k-num, 0) + slices.Min(dp[i:i+x])
	}
	return int64(slices.Min(dp[n:]))
}

// 选与不选
//
// - 时间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
// - 空间复杂度：$\mathcal{O}(n \cdot Σ)$。$Σ = 3$。
func minIncrementOperations3(nums []int, k int) int64 {
	var n = len(nums)

	// dfs(i ,j) 表示 nums[i] 的右边最近的大于等于 k 的数的距离，1<= j <= 3
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return 0
		}

		res = max(0, k-nums[i]) + dfs(i-1, 1)
		if j == Len {
			res = res // 距离达到最大值 3，只能选择 nums[i] 变成 k
		} else {
			res = min( // 距离未达到最大值 3，可以选与不选
				res,           // 选择 nums[i] 变成 k
				dfs(i-1, j+1), // 不选择 nums[i] 变成 k
			)
		}
		return res
	}

	cache := MemorySearch2(&dfs)
	_ = cache

	ans := dfs(n-1, 1)
	return int64(ans)
}

func minIncrementOperations4(nums []int, k int) int64 {
	x := 3

	n := len(nums)
	dp := make([]int, n+x)
	for i, num := range nums {
		dp[i+x] = max(k-num, 0) + slices.Min(dp[i:i+x])
	}
	//return int64(slices.Min(dp[n:])) // O(x*n) 的做法：

	return monotoneQueue(nums, k, x)
}

// 将
// 「任何长度 大于或等于 3 的子数组」
// 修改为：
// 「任何长度 大于或等于 x 的子数组 (1 <= x <= 1e5)」
//
// 通过单调队列优化，时间复杂度可以压缩为 O(n)
func monotoneQueue(nums []int, k, x int) int64 {
	var n = len(nums)
	dp := make([]int, n+1)
	queue := []int{0} // 哨兵，假设 nums[-1] = k，无需变大，因此 dp[-1] = 0。为方便计算，向右偏移一位
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for i-queue[0] > x {
			queue = queue[1:] // 去除过期元素
		}

		dp[i] = dp[queue[0]] + max(0, k-nums[i-1]) // 队首就是 [i-k:i-1] 中的最小值，对应着 Min(dp[max(i-k, 0):i])
		m := len(queue)
		for m > 0 && dp[i] <= dp[queue[m-1]] { // 单调递增队列
			m--
		}
		queue = append(queue[:m], i) // 新元素新队，维护单调性
	}
	return int64(slices.Min(dp[max(0, n-x+1):]))
}
