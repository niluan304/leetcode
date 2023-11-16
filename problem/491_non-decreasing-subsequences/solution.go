package main

import (
	"strconv"
)

func findSubsequences(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	var sum = 0
	var set = map[int]bool{
		0: true,
	}

	var dfs func(idx int)
	dfs = func(idx int) {
		if len(path) > 1 && !set[sum] {
			ans = append(ans, append([]int{}, path...))
			set[sum] = true
		}

		for i := idx; i < len(nums); i++ {
			v := nums[i]

			if len(path) > 1 && v < path[len(path)-1] {
				continue
			}

			sum = sum*10 + v
			path = append(path, v)
			dfs(i + 1)

			path = path[:len(path)-1]
			sum = sum / 10
		}
	}

	dfs(0)
	return ans
}

func findSubsequences2(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	var sum = ""
	var set = map[string]bool{}
	var dfs func(list []int)
	dfs = func(list []int) {
		for i, item := range list {
			if len(path) > 0 && item < path[len(path)-1] {
				continue
			}

			path = append(path, item)

			origin := len(sum)
			sum += strconv.Itoa(item) + "."

			if len(path) > 1 && !set[sum] {
				ans = append(ans, append([]int{}, path...))
				set[sum] = true
			}

			dfs(list[i+1:])
			path = path[:len(path)-1]
			sum = sum[:origin]
		}
	}

	for i := 0; i < len(nums); i++ {
		path = []int{}
		dfs(nums[i:])
	}
	return ans
}

func findSubsequences3(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	var sum = ""
	var set = map[string]bool{}
	var dfs func(list []int)
	dfs = func(list []int) {
		for i, item := range list {
			if len(path) > 0 && item < path[len(path)-1] {
				continue
			}

			path = append(path, item)

			origin := len(sum)
			sum += strconv.Itoa(item)

			if len(path) > 1 && !set[sum] {
				ans = append(ans, append([]int{}, path...))
				set[sum] = true
			}

			dfs(list[i+1:])
			path = path[:len(path)-1]
			sum = sum[:origin]
		}
	}

	for i := 0; i < len(nums); i++ {
		path = []int{}
		dfs(nums[i:])
	}
	return ans
}
