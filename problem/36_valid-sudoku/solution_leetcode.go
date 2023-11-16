package main

func leetcode(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
