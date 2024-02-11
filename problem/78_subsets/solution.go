package main

import (
	"slices"
)

// 子集型回溯模板1
//
// 从输入的角度出发（选与不选）
//
// - 时间复杂度：$\mathcal{O}(n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func subsets(nums []int) [][]int {
	var n = len(nums)
	var ans [][]int
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		dfs(i + 1) // 不选

		path = append(path, nums[i])
		dfs(i + 1)                // 选
		path = path[:len(path)-1] // 恢复现场
	}

	dfs(0)
	return ans
}

// 子集型回溯模板2
//
// 从答案的角度出发（选哪个数）
//
// - 时间复杂度：$\mathcal{O}(n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func subsets2(nums []int) [][]int {
	var n = len(nums)
	var ans [][]int
	var path []int
	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, slices.Clone(path))
		if i == n {
			return // 可以省去，因为下面的代码不会执行
		}
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}
