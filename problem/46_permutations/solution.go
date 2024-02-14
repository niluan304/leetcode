package main

import "slices"

func permute(nums []int) [][]int {
	var ans [][]int
	var path []int

	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(nums) == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		for i, num := range nums {
			path = append(path, num)
			x := append(slices.Clone(nums[:i]), nums[i+1:]...) // go1.22 slices.Concat
			dfs(x)
			path = path[:len(path)-1]
		}
	}

	dfs(nums)

	return ans
}

// 排列型回溯
// - 时间复杂度：$\mathcal{O}(n \cdot n!)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func permute2(nums []int) [][]int {
	var n = len(nums)
	var ans [][]int
	path := make([]int, n)
	onPath := make([]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		for j := 0; j < n; j++ {
			if onPath[j] {
				continue
			}
			path[i] = nums[j]
			onPath[j] = true
			dfs(i + 1)
			onPath[j] = false
		}
	}
	dfs(0)
	return ans
}
