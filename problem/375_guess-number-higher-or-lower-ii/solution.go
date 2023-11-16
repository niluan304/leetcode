package main

import (
	"math"
)

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

func toPtr(x int) *int { return &x }

// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2)
func getMoneyAmount(n int) int {
	var cache = make([][]*int, n+1)
	for i := 1; i <= n; i++ {
		cache[i] = make([]*int, n+2) // for i+1
		cache[i][i] = new(int)       //if i == j { return 0 }
		cache[i][i+1] = toPtr(i)     //if i+1 == j { return i }
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if cache[i][j] != nil {
			return *cache[i][j]
		}

		v := math.MaxInt32
		for k := i + 1; k <= j-1; k++ {
			v = _min(
				v,
				_max(
					dfs(i, k-1)+k, // 猜大了，往 [i, k-1] 区间猜
					//0,           // 猜中了
					dfs(k+1, j)+k, // 猜小了，往 [k+1, j] 区间猜
				),
			)
		}
		cache[i][j] = &v
		return v
	}

	return dfs(1, n)
}

// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2)
func getMoneyAmount2(n int) int {
	var cache = make([][]*int, n+1)
	for i := 1; i <= n; i++ {
		cache[i] = make([]*int, n+1)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}

		if cache[i][j] != nil {
			return *cache[i][j]
		}

		v := math.MaxInt32
		for k := i; k <= j; k++ {
			v = _min(v,
				_max(
					dfs(i, k-1)+k, // 猜大了，往 [i, k-1] 区间猜
					//0,           // 猜中了
					dfs(k+1, j)+k, // 猜小了，往 [k+1, j] 区间猜
				),
			)
		}
		cache[i][j] = &v
		return v
	}

	return dfs(1, n)
}

// dp 动态规划, 区间dp
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^2)
func getMoneyAmount3(n int) int {
	var dp = make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			v := math.MaxInt32
			for k := i; k <= j; k++ {
				v = _min(v,
					_max(
						dp[i][k-1]+k, // 猜大了，往 [i, k-1] 区间猜
						// 0          // 猜中了
						dp[k+1][j]+k, // 猜小了，往 [k+1, j] 区间猜
					),
				)
			}
			dp[i][j] = v
		}
	}

	return dp[1][n]
}
