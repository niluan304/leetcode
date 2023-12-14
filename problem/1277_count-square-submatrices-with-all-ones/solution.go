package main

func countSquares(matrix [][]int) int {
	var ans = 0
	s := NewMatrixSum(matrix)
	m, n := len(matrix), len(matrix[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			for r1, c1, i := r, c, 1; r1 >= 0 && c1 >= 0 && s.query2(r1, c1, r, c) == i*i; r1, c1, i = r1-1, c1-1, i+1 {
				ans++
			}
		}
	}
	return ans
}

// MatrixSum
// [【图解】二维前缀和（附模板代码 Python/Java/C++/Go/JS）(https://leetcode.cn/circle/discuss/UUuRex/)
// 作者：灵茶山艾府
type MatrixSum [][]int

func NewMatrixSum(matrix [][]int) MatrixSum {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range matrix {
		sum[i+1] = make([]int, n+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
		}
	}
	return sum
}

// 返回左上角在 (r1,c1) 右下角在 (r2-1,c2-1) 的子矩阵元素和（类似前缀和的左闭右开）
func (s MatrixSum) query(r1, c1, r2, c2 int) int {
	return s[r2][c2] - s[r2][c1] - s[r1][c2] + s[r1][c1]
}

// 如果你不习惯左闭右开，也可以这样写
// 返回左上角在 (r1,c1) 右下角在 (r2,c2) 的子矩阵元素和
func (s MatrixSum) query2(r1, c1, r2, c2 int) int {
	return s[r2+1][c2+1] - s[r2+1][c1] - s[r1][c2+1] + s[r1][c1]
}
