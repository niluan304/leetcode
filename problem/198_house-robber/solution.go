package main

import (
	. "github.com/niluan304/leetcode/copypasta/dp"
)

// 递归
// 未记忆化时的时间复杂度：O(2^n)
//
//	记忆化后的时间复杂度：O(n)
//
// 空间复杂度：O(n)
func rob(nums []int) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		return max(dfs(i-1), dfs(i-2)+nums[i])
	}

	MemorySearch(&dfs)
	return dfs(len(nums) - 1)
}

// 递归+缓存 = 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func rob2(nums []int) int {
	var dfs func(i int) int
	// 缓存会多次用到的 dfs(i)
	var cache = map[int]int{}
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		// 缓存之后，就不会不需要继续 "递" 了
		if v, ok := cache[i]; ok {
			return v
		}

		v := max(dfs(i-1), dfs(i-2)+nums[i])
		cache[i] = v
		return v
	}

	return dfs(len(nums) - 1)
}

// 动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func rob3(nums []int) int {
	var n = len(nums)
	// 1. dp数组的含义
	// dp[i+2] 表示第i个房间能拿到的最高金额
	var dp = make([]int, n+1)

	// 3. dp数组初始化：
	//dp[0] = 0
	dp[1] = nums[0]

	// 4. 遍历顺序
	// 从底到顶
	for i := 1; i < n; i++ {
		// 2. 第i个房间的最大值, 选不选 i 的问题
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}

	//// 5. debug 打印dp数组
	//fmt.Println("dp", dp)
	return dp[n]
}

// 动态规划, 空间优化
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func rob4(nums []int) int {
	var n = len(nums)
	var dp0, dp1 = 0, 0

	for i := 0; i < n; i++ {
		// 2. 第i个房间的最大值, 选不选 i 的问题
		v := max(dp0, dp1+nums[i])
		dp1 = dp0
		dp0 = v

		//// 5. debug 打印dp数组
		//fmt.Print(dp0, " ")
	}
	return dp0
}
