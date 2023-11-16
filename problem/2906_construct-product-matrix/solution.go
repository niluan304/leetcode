package main

const mod = 12345

func constructProductMatrix(grid [][]int) [][]int {
	var n = len(grid)
	var m = len(grid[0])

	var prefix, suffix = make([]int, m*n+1), make([]int, m*n+1)

	prefix[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			k := i*m + j
			prefix[k+1] = prefix[k] * grid[i][j] % mod
		}
	}

	suffix[n*m] = 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			k := i*m + j
			suffix[k] = suffix[k+1] * grid[i][j] % mod
		}
	}

	var ans = make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			k := i*m + j
			ans[i][j] = prefix[k] * suffix[k+1] % mod
		}
	}
	return ans
}

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
