package main

import (
	"sort"
)

const MOD = 1_000_000_007

func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	var n = len(nums)

	var dp, pow, ans = make([]int, n), make([]int, n), 0
	pow[0] = 1
	for i := 1; i < n; i++ {
		pow[i] = 2 * pow[i-1] % MOD
		x := (nums[i] - nums[i-1]) * (pow[i] - 1)

		dp[i] = (x%MOD + 2*dp[i-1]%MOD) % MOD
		ans = (ans + dp[i]) % MOD
	}
	return ans
}

func sumSubseqWidths2(nums []int) int {
	sort.Ints(nums)
	var n = len(nums)

	var dp, pow, ans = 0, 1, 0
	for i := 1; i < n; i++ {
		pow = 2 * pow % MOD
		x := (nums[i] - nums[i-1]) * (pow - 1)

		dp = (x%MOD + 2*dp%MOD) % MOD
		ans = (ans + dp) % MOD
	}
	return ans
}
