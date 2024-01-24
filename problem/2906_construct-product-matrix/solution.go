package main

const MOD = 12345

// 前后缀分解
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
func constructProductMatrix(grid [][]int) [][]int {
	var n, m = len(grid), len(grid[0])
	mx := n * m
	var prefix, suffix = make([]int, mx+1), make([]int, mx+1)
	prefix[0], suffix[mx] = 1, 1 // 乘法初始化为 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			prefix[idx+1] = (prefix[idx] % MOD) * (grid[i][j] % MOD) % MOD
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			idx := i*m + j
			suffix[idx] = (suffix[idx+1] % MOD) * (grid[i][j] % MOD) % MOD
		}
	}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			idx := i*m + j
			grid[i][j] = (prefix[idx] * suffix[idx+1]) % MOD
		}
	}
	return grid
}
