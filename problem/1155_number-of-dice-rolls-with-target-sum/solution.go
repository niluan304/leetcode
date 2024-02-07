package main

func numRollsToTarget(n int, k int, target int) int {
	const mod = 1e9 + 7
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j <= target; j++ {
			v := 0
			for x := 1; x <= k && x <= j; x++ {
				v = (v + dp[i-1][j-x]) % mod
			}
			dp[i][j] = v
		}
	}
	return dp[n][target]
}

func numRollsToTarget2(n int, k int, target int) int {
	const mod = 1e9 + 7
	var cache = make([][]*int, n+1)
	for i, _ := range cache {
		cache[i] = make([]*int, target+1)
	}

	var dfs func(n, target int) int
	dfs = func(n, target int) int {
		if n == 0 {
			if target == 0 {
				return 1
			}
			return 0
		}

		ptr := &cache[n][target] // 取指针的指针
		if *ptr != nil {
			return **ptr // 计算过
		}
		*ptr = new(int) // 初始化 cache[n][target]，等价于 cache[n][target] = new(int); ptr = &cache[n][target];

		for x := 1; x <= k && x <= target; x++ {
			**ptr = (**ptr + dfs(n-1, target-x)) % mod
		}
		return **ptr
	}

	return dfs(n, target)
}
