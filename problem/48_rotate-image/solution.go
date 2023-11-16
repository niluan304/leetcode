package main

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	var m = make([][]int, len(matrix))

	for i := range m {
		m[i] = make([]int, len(matrix))
	}

	for i, row := range matrix {
		for j, column := range row {
			m[j][n-i] = column
		}
	}

	for i, row := range m {
		for j := range row {
			matrix[i][j] = m[i][j]
		}
	}
}

// rotate2 未通过 Case2,Case3 测试
func rotate2(matrix [][]int) {
	n := len(matrix) - 1
	y0, x0 := len(matrix)/2-1, len(matrix)/2+1

	for y := y0; y >= 0; y-- {
		for x := y; x < x0; x++ {
			matrix[x][y], matrix[y][n-x], matrix[n-x][n-y], matrix[n-y][x] =
				matrix[n-y][x], matrix[x][y], matrix[y][n-x], matrix[n-x][n-y]
		}
	}
}

// rotate3 逐层旋转
// 第一次: 旋转 [1, n-1]
// 第二次: 旋转 [2, n-2]
// 第三次: 旋转 [3, n-3]
func rotate3(matrix [][]int) {
	var (
		l = len(matrix)
		n = l - 1
	)

	for y := 0; y < l/2+1; y++ {
		for x := y; x < n-y; x++ {
			matrix[x][y], matrix[y][n-x], matrix[n-x][n-y], matrix[n-y][x] =
				matrix[n-y][x], matrix[x][y], matrix[y][n-x], matrix[n-x][n-y]
		}
	}
}

func leetcode1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// leetcode 2 先水平翻转, 再主对角线翻转
func leetcode2(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
