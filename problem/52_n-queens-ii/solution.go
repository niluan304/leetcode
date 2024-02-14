package main

func totalNQueens(n int) int {
	m := 2*n - 1

	col := make([]int, n)     // 在第 i 行 col[i] 列放置皇后
	onPath := make([]bool, n) // 哪些行选了皇后
	diag1, diag2 := make(map[int]bool, m), make(map[int]bool, m)

	var ans = 0
	var dfs func(r int) // 枚举行号
	dfs = func(r int) {
		if r == n {
			ans++
		}

		for c := 0; c < n; c++ {
			if onPath[c] || diag1[r+c] || diag2[r-c] {
				continue
			}
			col[r] = c
			onPath[c], diag1[r+c], diag2[r-c] = true, true, true
			dfs(r + 1)
			onPath[c], diag1[r+c], diag2[r-c] = false, false, false
		}
	}
	dfs(0)
	return ans
}
