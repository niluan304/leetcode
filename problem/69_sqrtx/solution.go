package main

import (
	"math"
	"sort"
)

func mySqrt(x int) int {
	var ans = -1
	var left, right = 0, x
	for left <= right {
		mid := (right-left)>>1 + left

		y := mid * mid
		if y > x {
			right = mid - 1
		} else {
			left = mid + 1
			ans = mid
		}
	}

	return ans
}

func stdlib(x int) int {
	return int(math.Sqrt(float64(x)))
}

func mySqrt2(x int) int {
	var left, right = 0, x

	for left <= right {
		mid := (right-left)>>1 + left

		y := mid * mid
		if y > x {
			right = mid - 1
		} else if y < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}

func mySqrt3(x int) int {
	return sort.Search(x+1, func(i int) bool { return i*i > x }) - 1
}
