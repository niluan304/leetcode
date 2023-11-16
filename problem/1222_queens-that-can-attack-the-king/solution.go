package main

import "image"

// 模拟：寻找国王的8个方向上是否有皇后
// 时间复杂度：O(n+8)
// 时间复杂度：O(8*8)
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	// grid 网格
	// matrix 矩阵
	// chessboard 棋盘

	const N = 8
	var findQueen = [N][N]bool{}
	for _, queen := range queens {
		x, y := queen[0], queen[1]
		findQueen[x][y] = true
	}

	moves := []image.Point{
		{X: 0, Y: +1},
		{X: 0, Y: -1},
		{X: +1, Y: 0},
		{X: -1, Y: 0},
		{X: +1, Y: +1},
		{X: +1, Y: -1},
		{X: -1, Y: +1},
		{X: -1, Y: -1},
	}

	var ans = make([][]int, 0, N*N)
	for _, move := range moves {
		x, y := king[0], king[1]
		// 寻找该方向上是否有皇后
		for x >= 0 && y >= 0 && x < N && y < N {
			// 首次为国王坐标, 必然 false
			if findQueen[x][y] {
				ans = append(ans, []int{x, y})
				break
			}
			x += move.X
			y += move.Y
		}
	}
	return ans
}
