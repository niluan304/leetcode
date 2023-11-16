package main

func setZeroes(matrix [][]int) {
	var rows, columns = make([]bool, len(matrix)), make([]bool, len(matrix[0]))

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for i, row := range matrix {
		for j := range row {
			if rows[i] || columns[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// setZeroes2 移位会产生溢出: 1 <= m, n <= 200
func setZeroes2(matrix [][]int) {
	var rows, columns = 0, 0

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				rows |= 1 << i
				columns |= 1 << j
			}
		}
	}

	for i, row := range matrix {
		for j := range row {
			if rows&(1<<i) > 0 || columns&(1<<j) > 0 {
				matrix[i][j] = 0
			}
		}
	}
}

func leetcode2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}

func leetcode3(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
