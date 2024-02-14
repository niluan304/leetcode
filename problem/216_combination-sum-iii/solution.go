package main

import "slices"

// 组合型回溯模板1
// 从输入的角度出发：
//
// 回溯三问：
// 当前操作？
// 子问题？
// 下一个子问题？
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int

	// 2 <= k <= 9
	// 1 <= n <= 60
	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		if len(path) == k {
			if sum == n {
				ans = append(ans, slices.Clone(path))
			}
			return
		}
		if i == 10 || sum > n { // 不合法的方案
			return
		}

		dfs(i+1, sum) // 不选

		path = append(path, i) // 选
		dfs(i+1, sum+i)
		path = path[:len(path)-1]
	}
	dfs(1, 0)
	return ans
}

// 组合型回溯模板2
// 从答案的角度出发：
//
// 回溯三问：
// 当前操作？ 选择 i 来拼凑 n
// 子问题？ 从 [i, 9] 的 k 个数中拼凑出 n
// 下一个子问题？ 从 [i+1, 9] 的 k-1 个数中 拼凑出 n-i
func combinationSum3ii(k int, n int) [][]int {
	var ans [][]int
	var path []int

	// 2 <= k <= 9
	// 1 <= n <= 60
	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		if sum > n {
			return
		}

		if len(path) == k {
			if sum == n {
				ans = append(ans, slices.Clone(path))
			}
			return
		}
		for j := i; j <= 9; j++ {
			path = append(path, j)
			dfs(j+1, sum+j)
			path = path[:len(path)-1]
		}
	}
	dfs(1, 0)
	return ans
}
