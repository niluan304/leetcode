package main

import . "github.com/niluan304/leetcode/copypasta/dp"

// dfs + 记忆化搜索
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func validPartition(nums []int) bool {
	n := len(nums)

	// dfs(i) 求 nums[i:] 是否为合法划分
	var dfs func(i int) bool
	dfs = func(i int) bool {
		// 递归边界，表示空数组 nums[n:] 是合法的划分
		if i == n {
			return true
		}
		//// 如果数组长度不够, c1 c2 c3 会是 false，所以无须担心数组越界
		//if i > n {
		//	return false
		//}

		x := nums[i]

		c1 := i+1 < n && (x == nums[i+1])                       // 2 个相等元素的子数组
		c2 := i+2 < n && (x == nums[i+1] && x == nums[i+2])     // 3 个相等元素的子数组
		c3 := i+2 < n && (x+1 == nums[i+1] && x+2 == nums[i+2]) // 公差为 +1 的等差子数组

		return (c1 && dfs(i+2)) || (c2 && dfs(i+3)) || (c3 && dfs(i+3))
	}

	MemorySearch(&dfs)

	ans := dfs(0)
	return ans
}

// 将 validPartition 翻译为 动态规划
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func validPartition2(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1) // dp[i+1] 求 nums[:i] 是否为合法划分
	dp[0] = true            // 初始化，表示空数组 nums[:0] 是合法的划分

	for i, x := range nums {
		c1 := i >= 1 && (x == nums[i-1])                       // 2 个相等元素的子数组
		c2 := i >= 2 && (x == nums[i-1] && x == nums[i-2])     // 3 个相等元素的子数组
		c3 := i >= 2 && (x-1 == nums[i-1] && x-2 == nums[i-2]) // 公差为 +1 的等差子数组

		dp[i+1] = (c1 && dp[i-1]) || (c2 && dp[i-2]) || (c3 && dp[i-2])
	}
	return dp[n]
}
