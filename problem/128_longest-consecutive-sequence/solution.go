package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	var max = 1
	var cur = 1

	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 0:
			continue
		case 1:
			cur++
		default:
			cur = 1
		}

		if max < cur {
			max = cur
		}
	}

	return max
}
