package main

import (
	"math"
)

// 题目输入不包含0, 不需要考虑0为最小值的情况
func minNumber(nums1 []int, nums2 []int) int {
	var common = make([]bool, 10)
	var x1, x2 = math.MaxInt, math.MaxInt
	for _, v := range nums1 {
		common[v] = true
		x1 = _min(x1, v)
	}

	var ans = math.MaxInt
	for _, v := range nums2 {
		if common[v] {
			ans = _min(ans, v)
		}
		x2 = _min(x2, v)
	}

	if ans != math.MaxInt {
		return ans
	}

	// 题目输入不包含0, 不需要考虑0为最小值的情况
	if x1 > x2 {
		x2, x1 = x1, x2
	}
	return x1*10 + x2
}

func _min(list ...int) int {
	var ans = math.MaxInt
	for _, n := range list {
		if ans > n {
			ans = n
		}
	}

	return ans
}
