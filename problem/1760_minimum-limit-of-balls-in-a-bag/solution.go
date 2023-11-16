package main

import (
	"sort"
)

func minimumSize(nums []int, maxOperations int) int {
	max, min := 0, 1
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	var size = 0
	for min <= max {
		var (
			opt = 0
			mid = (max + min) / 2
		)

		for _, num := range nums {
			opt += (num - 1) / mid
		}

		if opt <= maxOperations {
			size = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return size
}

func leetcode(nums []int, maxOperations int) int {
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}
	return sort.Search(max, func(y int) bool {
		if y == 0 {
			return false
		}
		ops := 0
		for _, x := range nums {
			ops += (x - 1) / y
		}
		return ops <= maxOperations
	})
}
