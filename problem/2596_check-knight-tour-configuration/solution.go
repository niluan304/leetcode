package main

import "image"

// 模拟巡视, 只要每一步都是合法的, 那么就可以完成巡视
// 时间复杂度：O(8*n)
// 空间复杂度：O(n)
func checkValidGrid(grid [][]int) bool {
	var n = len(grid)

	// 在有效的巡视方案中，骑士会从棋盘的 左上角 出发，并且访问棋盘上的每个格子 恰好一次 。
	// 但题目没说 grid[0][0] 一定是 0
	if grid[0][0] != 0 {
		return false
	}

	var nextPoints = func(x, y int) (points []image.Point) {
		return []image.Point{
			{X: x + 2, Y: y + 1},
			{X: x + 2, Y: y - 1},
			{X: x - 2, Y: y + 1},
			{X: x - 2, Y: y - 1},
			{X: x + 1, Y: y + 2},
			{X: x + 1, Y: y - 2},
			{X: x - 1, Y: y + 2},
			{X: x - 1, Y: y - 2},
		}
	}

	var x, y = 0, 0
Next:
	for i := 1; i < n*n; i++ {
		points := nextPoints(x, y)
		for _, point := range points {
			x, y = point.X, point.Y
			if x < 0 || y < 0 || x >= n || y >= n {
				continue
			}

			// 下一步是对的坐标，继续巡视
			if grid[x][y] == i {
				continue Next
			}
		}
		return false
	}

	return true
}
