package main

import (
	"strconv"
)

func findSubsequences(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	sum := ""
	set := map[string]bool{}
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

func findSubsequences2(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	sum := ""
	set := map[string]bool{}
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
