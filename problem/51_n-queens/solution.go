package main

import (
	"bytes"
	"strings"
)

func solveNQueens(n int) [][]string {
	m := 2*n - 1

	col := make([]int, n)     // 在第 i 行 col[i] 列放置皇后
	onPath := make([]bool, n) // 哪些行选了皇后
	diag1, diag2 := make(map[int]bool, m), make(map[int]bool, m)

	var ans [][]string
	var dfs func(r int) // 枚举行号
	dfs = func(r int) {
		if r == n {
			var path = make([]string, n)
			for j, c := range col {
				row := strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
				path[j] = row
			}
			ans = append(ans, path)
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

func solveNQueens2(n int) [][]string {
	var ans [][]string
	const Empty, Queen = '.', 'Q'
	var board = make([][]byte, n)
	for i := range board {
		board[i] = bytes.Repeat([]byte{Empty}, n)
	}

	valid := func(board [][]byte, x, y int) bool {
		for i, v := range board[:x] {
			if v[y] == Queen {
				return false
			}

			y1 := y - (x - i)
			if y1 >= 0 && v[y1] == Queen {
				return false
			}

			y2 := y + (x - i)
			if y2 < n && v[y2] == Queen {
				return false
			}
		}
		return true
	}

	var dfs func(idx int)
	dfs = func(x int) {
		if x == n {
			var path = make([]string, n)
			for i := range board {
				path[i] = string(board[i])
			}
			ans = append(ans, path)
			return
		}

		for y := range board[x] {
			if !valid(board, x, y) {
				continue
			}

			board[x][y] = Queen
			dfs(x + 1)
			board[x][y] = Empty
		}
	}

	dfs(0)
	return ans
}
