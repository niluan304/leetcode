package main

// dfs + 记忆化搜索
//
// - 时间复杂度：O(mn)。
// - 空间复杂度：O(mn)。
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	cache := make([][]*int, m)
	for i := range cache {
		cache[i] = make([]*int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if j == n-1 {
			return 0
		}

		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		// 向右上/右/右下走一步
		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				// 其实 res 的值，就等于最深的 j，可以不用设置返回值，转换为 ans = max(ans, j)
				res = max(res, dfs(k, j+1)+1)
			}
		}
		return res
	}

	ans := 0
	for i := 0; i < m; i++ {
		ans = max(ans, dfs(i, 0))
	}
	return ans
}

// dfs
//
// maxMoves 的优化
// 通过操作 grid[i][j] = 0，替代了记忆化搜索
//
// - 时间复杂度：O(mn)。
// - 空间复杂度：O(mn)。
func maxMoves2(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	ans := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		ans = max(ans, j)

		if j == n-1 {
			return
		}

		// 向右上/右/右下走一步
		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				// 其实 res 的值，就等于最深的 j，可以不用设置返回值，转换为 ans = max(ans, j)
				dfs(k, j+1)
			}
		}

		grid[i][j] = 0
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
	}
	return ans
}
