package main

import (
	. "github.com/niluan304/leetcode/copypasta/dp"
)

const MOD = 1_000_000_007

func numberOfWays(n int, x int) int {
	type Pair struct {
		Target int
		N      int
	}

	var cache = make(map[Pair]int, n)
	var dfs func(target, n int) int
	dfs = func(target, n int) (res int) {
		key := Pair{Target: target, N: n}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		if target == 0 {
			return 1
		}

		pow := Pow(n, x)
		if pow > target {
			return 0
		}

		res = (dfs(target-pow, n+1) + dfs(target, n+1)) % MOD
		return res
	}

	ans := dfs(n, 1)
	return ans
}

func Pow(x int, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / Pow(x, -n)
	}

	if n%2 == 0 {
		return Square(Pow(x, n/2))
	}

	return Square(Pow(x, (n-1)/2)) * x
}

func Square(x int) int { return x * x }

func numberOfWays2(n int, x int) int {
	var nums = make([]int, 0, n)
	for i := 1; i <= n; i++ {
		pow := Pow(i, x)
		if pow > n {
			break
		}
		nums = append(nums, pow)
	}

	return ZeroOneWaysToSum(nums, n) % MOD
}

func numberOfWays3(n int, x int) int {
	sum := n
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, sum+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		num := Pow(i+1, x) // nums = [1^x, 2^x,...,n^x], nums[i] = (i+1)^x

		for s := 0; s <= sum; s++ {
			dp[i+1][s] = dp[i][s]
			if s >= num {
				dp[i+1][s] += dp[i][s-num] // % mod
			}
		}
	}
	return dp[n][n] % MOD
}
