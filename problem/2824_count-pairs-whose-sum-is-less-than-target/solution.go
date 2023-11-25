package main

import "sort"

func countPairs(nums []int, target int) int {
	var ans = 0

	for i, num := range nums {
		for _, x := range nums[i+1:] {
			if num+x < target {
				ans++
			}
		}
	}
	return ans
}

func countPairs2(nums []int, target int) int {
	var ans = 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		ans += sort.SearchInts(nums[:i], target-nums[i])
	}
	return ans
}
