package main

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	var set = map[int]bool{}

	for i := range nums {
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	var set = map[int]struct{}{}

	for i := range nums {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
	}

	return false
}

func containsDuplicate3(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func leetcode1(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
