package main

import "strings"

func leetcodeCarl(n int) [][]string {
	var res [][]string
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtrack(row + 1)
				chessboard[row][i] = "."
			}
		}
	}
	backtrack(0)
	return res
}

func isValid(n, row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}

// 执行用时为 0 ms 的范例
func leetcodeMinTime(n int) [][]string {
	var solveNQueensResults [][]string

	square := make([]string, n)
	for i := range square {
		square[i] = strings.Repeat(".", n)
	}

	var backtrack func(square []string, row, n int)
	backtrack = func(square []string, row, n int) {
		if len(square) == row {
			tmp := make([]string, len(square))
			copy(tmp, square)
			solveNQueensResults = append(solveNQueensResults, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if isValidMinTime(square, n, row, col) {
				r := []byte(square[row])
				r[col] = 'Q'
				square[row] = string(r)
				backtrack(square, row+1, n)
				r[col] = '.'
				square[row] = string(r)
			}
		}
	}
	backtrack(square, 0, n)
	return solveNQueensResults
}

func isValidMinTime(square []string, n, row, col int) bool {
	for i := 0; i < n; i++ {
		if i != row && square[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if square[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	for i, j := row-1, col+1; i >= 0 && j < n; {
		if square[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}

// 执行消耗内存为 2924 kb 的范例
func leetcodeMinMemory(n int) (res [][]string) {
	placements := make([]string, n)
	for i := range placements {
		buf := make([]byte, n)
		for j := range placements {
			if i == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		placements[i] = string(buf)
	}
	var construct func(prev []int)
	construct = func(prev []int) {
		if len(prev) == n {
			plan := make([]string, n)
			for i := 0; i < n; i++ {
				plan[i] = placements[prev[i]]
			}
			res = append(res, plan)
			return
		}
		occupied := 0
		for i := range prev {
			dist := len(prev) - i
			bit := 1 << prev[i]
			occupied |= bit | bit<<dist | bit>>dist
		}
		prev = append(prev, -1)
		for i := 0; i < n; i++ {
			if (occupied>>i)&1 != 0 {
				continue
			}
			prev[len(prev)-1] = i
			construct(prev)
		}
	}
	construct(make([]int, 0, n))
	return
}
