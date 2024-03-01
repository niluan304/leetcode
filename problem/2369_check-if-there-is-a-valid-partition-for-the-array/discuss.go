package main

import . "github.com/niluan304/leetcode/copypasta/dp"

// TODO
// 思考题
// 有多少种有效划分的方案数？模 1e9+7

// dfs + 记忆化搜索
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func discuss(nums []int) int {
	const Mod = 1e9 + 7

	n := len(nums)
	var dfs func(i int) (res int)
	dfs = func(i int) (res int) {
		if i == n {
			return 1
		}

		x := nums[i]

		c1 := i+1 < n && (x == nums[i+1])                       // 2 个相等元素的子数组
		c2 := i+2 < n && (x == nums[i+1] && x == nums[i+2])     // 3 个相等元素的子数组
		c3 := i+2 < n && (x+1 == nums[i+1] && x+2 == nums[i+2]) // 公差为 +1 的等差子数组

		if c1 {
			res += dfs(i + 2)
		}
		if c2 || c3 {
			res += dfs(i + 3)
		}
		return res % Mod
	}

	MemorySearch(&dfs)

	ans := dfs(0)
	return ans
}
