package main

import "math"

func reverse(x int) int {
	var positive = x > 0 // 是否为正数

	x = max(x, -x)

	var ans = 0
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if !positive {
		ans = -ans
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return ans
}
