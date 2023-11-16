package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	sort.Ints(candidates)

	var dfs func(idx int, sum int)
	dfs = func(idx int, sum int) {
		if sum >= target {
			if sum == target {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}

		for _, v := range candidates[idx:] {
			if len(path) > 0 && path[len(path)-1] > v {
				continue
			}

			path = append(path, v)
			dfs(idx, sum+v)
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	// 做升序排序, 方便剪枝
	sort.Ints(candidates)

	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := idx; i < len(candidates); i++ {
			v := candidates[i]

			// candidates是升序的，如果target < v，后面的元素肯定都不满足
			if target < v {
				return
			}
			path = append(path, v)

			// idx+i: 从当前元素 v 开始搜索
			dfs(target-v, i)
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)
	return ans
}
