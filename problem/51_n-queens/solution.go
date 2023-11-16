package main

import (
	"bytes"
)

func solveNQueens(n int) [][]string {
	var (
		ans  [][]string
		path []string
	)
	q := bytes.Repeat([]byte("."), n)

	var matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	var dfs func(idx int)

	dfs = func(idx int) {
		if len(path) == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}

		for x := idx; x < n; x++ {
			for y, v := range matrix[x] {
				if v > 0 {
					continue
				}

				node := append([]byte{}, q...)
				node[y] = 'Q'
				path = append(path, string(node))

				for i := x; i < n; i++ {
					matrix[i][y]++

					y1 := y - (i - x)
					if y1 >= 0 {
						matrix[i][y1]++
					}

					y2 := y + (i - x)
					if y2 < n {
						matrix[i][y2]++
					}
				}

				dfs(x + 1)

				path = path[:len(path)-1]
				for i := x; i < n; i++ {
					matrix[i][y]--

					y1 := y - (i - x)
					if y1 >= 0 {
						matrix[i][y1]--
					}

					y2 := y + (i - x)
					if y2 < n {
						matrix[i][y2]--
					}
				}
			}
		}
	}

	dfs(0)

	return ans
}

func solveNQueens2(n int) [][]string {
	var (
		ans [][]string
	)

	var board = make([][]byte, n)
	for i := range board {
		board[i] = bytes.Repeat([]byte("."), n)
	}

	valid := func(x, y int) bool {
		for i, v := range board[:x] {
			if v[y] == 'Q' {
				return false
			}

			y1 := y - (x - i)
			if y1 >= 0 && v[y1] == 'Q' {
				return false
			}

			y2 := y + (x - i)
			if y2 < n && v[y2] == 'Q' {
				return false
			}
		}

		return true
	}

	var dfs func(idx int)

	dfs = func(idx int) {
		if idx == n {
			var path = make([]string, 0, n)
			for i := range board {
				path = append(path, string(board[i]))
			}

			ans = append(ans, path)
			return
		}

		for x := idx; x < n; x++ {
			flag := false
			for y := range board[x] {
				if !valid(x, y) {
					continue
				}

				flag = true
				board[x][y] = 'Q'

				dfs(x + 1)

				flag = false
				board[x][y] = '.'
			}

			// 未进入下一行, 说明这一行无法设置Q
			if !flag {
				return
			}
		}
	}

	dfs(0)
	return ans
}

func solveNQueens3(n int) [][]string {
	var (
		ans [][]string
	)

	var board = make([][]byte, n)
	for i := range board {
		board[i] = bytes.Repeat([]byte("."), n)
	}

	var dfs func(idx int)

	dfs = func(x int) {
		if x == n {
			var path = make([]string, 0, n)
			for i := range board {
				path = append(path, string(board[i]))
			}

			ans = append(ans, path)
			return
		}

		for y := range board[x] {
			if !IsValid3(board, x, y) {
				continue
			}

			board[x][y] = 'Q'
			dfs(x + 1)
			board[x][y] = '.'
		}
	}

	dfs(0)
	return ans
}

func IsValid3(board [][]byte, x, y int) bool {
	n := len(board)
	for i, v := range board[:x] {
		if v[y] == 'Q' {
			return false
		}

		y1 := y - (x - i)
		if y1 >= 0 && v[y1] == 'Q' {
			return false
		}

		y2 := y + (x - i)
		if y2 < n && v[y2] == 'Q' {
			return false
		}
	}

	return true
}
