package main

import (
	"sort"
)

const Mod int = 1e9 + 7

func sumDistance(nums []int, s string, d int) int {
	for i := range nums {
		switch s[i] {
		case 'L':
			nums[i] -= d
		case 'R':
			nums[i] += d
		}
	}

	sort.Ints(nums)

	var n = len(nums) - 1 // n 个间隔
	var ans = 0
	for i := 1; i <= n; i++ {
		v := (nums[i] - nums[i-1]) % Mod
		ans += v * i * (n - i + 1) % Mod
	}

	return ans % Mod
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
