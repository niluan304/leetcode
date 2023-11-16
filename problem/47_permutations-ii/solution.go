package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	var path []int

	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(nums) == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])

			list := append(append([]int{}, nums[:i]...), nums[i+1:]...)
			dfs(list)

			path = path[:len(path)-1]
		}
	}

	dfs(nums)
	return ans
}
