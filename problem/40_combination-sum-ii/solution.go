package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	type Value struct {
		Value int
		Count int
	}

	var values []Value
	// 将path转化为数值, 以方便重复判断,
	// 如果 path = [1,2,2,5], sum = x1225, 其中x为len(path)
	var counts = make(map[int]int, len(candidates))
	for _, candidate := range candidates {
		counts[candidate]++
	}

	for value, count := range counts {
		values = append(values, Value{
			Value: value,
			Count: count,
		})
	}

	// 升序排序, 方便剪枝
	sort.Slice(values, func(i, j int) bool {
		return values[i].Value < values[j].Value
	})

	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := idx; i < len(values); i++ {
			v := values[i].Value
			if target < v {
				return
			}

			if values[i].Count == 0 {
				continue
			}

			path = append(path, v)
			values[i].Count--

			dfs(target-v, i)
			path = path[:len(path)-1]
			values[i].Count++
		}
	}

	dfs(target, 0)
	return ans
}

func combinationSum2II(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	// 升序排序, 方便剪枝
	sort.Ints(candidates)

	type Value struct {
		Value int
		Count int
	}

	var values []Value
	for _, candidate := range candidates {
		n := len(values)

		if n == 0 || values[n-1].Value != candidate {
			values = append(values, Value{
				Value: candidate,
				Count: 1,
			})
		} else {
			values[n-1].Count++
		}
	}

	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := idx; i < len(values); i++ {
			v := values[i].Value
			if target < v {
				return
			}

			if values[i].Count == 0 {
				continue
			}

			path = append(path, v)
			values[i].Count--

			dfs(target-v, i)
			path = path[:len(path)-1]
			values[i].Count++
		}
	}

	dfs(target, 0)
	return ans
}

func combinationSum2III(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	// 升序排序, 方便剪枝
	sort.Ints(candidates)

	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := idx; i < len(candidates); i++ {
			v := candidates[i]
			if target < v {
				return
			}

			if i != idx && v == candidates[i-1] {
				continue
			}

			path = append(path, v)
			dfs(target-v, i+1)
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)
	return ans
}

func combinationSum2IV(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	// 将path转化为数值, 以方便重复判断,
	// 如果 path = [1,2,2,5], sum = x1225, 其中x为len(path)
	var set = make(map[int]struct{}, len(candidates))

	// 升序排序, 方便剪枝
	sort.Ints(candidates)

	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if target == 0 {
			n := len(path)
			sum := n * 10

			for _, item := range path {
				sum = sum*10 + item
			}

			if _, ok := set[sum]; ok {
				return
			}

			set[sum] = struct{}{}

			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := idx; i < len(candidates); i++ {
			v := candidates[i]
			if target < v {
				return
			}

			path = append(path, v)
			dfs(target-v, i+1)
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)
	return ans
}
