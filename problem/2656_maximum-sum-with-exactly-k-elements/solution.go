package main

import "slices"

func maximizeSum(nums []int, k int) int {
	num := slices.Max(nums)
	return (num + num + k - 1) * k / 2
}
