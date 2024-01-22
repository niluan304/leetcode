package main

import (
	"math"
)

func maximumJumps(nums []int, target int) int {

	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32

	var cache = make([]*int, n)
	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i == 0 {
			return 0
		}

		ptr := &cache[i]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		res = math.MinInt32
		for j := 0; j < i; j++ {
			if Abs(nums[j]-nums[i]) <= target {
				res = max(res, dfs(j)+1)
			}
		}
		return res
	}

	ans = dfs(n - 1)
	if ans < 0 {
		return -1
	}
	return ans
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
