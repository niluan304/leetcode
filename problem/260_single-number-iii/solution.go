package main

import "math/bits"

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	tz := bits.TrailingZeros(uint(xorSum))

	var ans = make([]int, 2)
	for _, num := range nums {
		ans[num>>tz&1] ^= num
	}
	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
