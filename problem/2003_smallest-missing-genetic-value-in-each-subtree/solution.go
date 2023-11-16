package main

import "slices"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	var n = len(parents)

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	node := slices.Index(nums, 1)
	if node < 0 { // 不存在基因值为 1 的点
		return ans
	}

	var children = make([][]int, n)
	for child := 1; child < n; child++ {
		parent := parents[child]
		children[parent] = append(children[parent], child)
	}

	visit := make(map[int]bool, n)
	var dfs func(node int)
	dfs = func(node int) {
		visit[nums[node]] = true

		for _, child := range children[node] {
			if !visit[nums[child]] {
				dfs(child)
			}
		}
	}

	for x := 2; node != -1; node = parents[node] {
		dfs(node)
		for visit[x] {
			x++
		}
		ans[node] = x
	}

	return ans
}

func smallestMissingValueSubtree2(parents []int, nums []int) []int {
	var n = len(parents)

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	node := slices.Index(nums, 1)
	if node < 0 { // 不存在基因值为 1 的点
		return ans
	}

	var children = make([][]int, n)
	for child := 1; child < n; child++ {
		parent := parents[child]
		children[parent] = append(children[parent], child)
	}

	visit := make(map[int]bool, n)
	var dfs func(node int)
	dfs = func(node int) {
		visit[nums[node]] = true

		for _, child := range children[node] {
			if !visit[nums[child]] {
				dfs(child)
			}
		}
	}

	for x := 2; node != -1; node = parents[node] {
		dfs(node)
		for visit[x] {
			x++
		}
		ans[node] = x
	}
	return ans
}
