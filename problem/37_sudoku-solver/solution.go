package main

const (
	Size  = 9
	Block = 3
)

func solveSudoku(board [][]byte) {
	var (
		rows    = [Size][Size]int{}
		columns = [Size][Size]int{}
		blocks  = [Size][Size]int{}
	)

	for x, row := range board {
		for y, b := range row {
			if b == '.' {
				continue
			}
			v := b - '1'

			n := (x / 3) + (y/3)*3
			rows[x][v] = 1
			columns[y][v] = 1
			blocks[n][v] = 1
		}
	}

	var dfs func(x, y int)

	found := false
	dfs = func(x, y int) {
		if y == Size {
			x++
			y = 0
			if x == Size {
				found = true
				return
			}
		}

		if board[x][y] != '.' {
			dfs(x, y+1)
			return
		}

		n := (x / 3) + (y/3)*3
		for _, v := range Values(rows[x], columns[y], blocks[n]) {
			rows[x][v] = 1
			columns[y][v] = 1
			blocks[n][v] = 1
			board[x][y] = byte(v) + '1'

			dfs(x, y+1)
			if found {
				return
			}

			rows[x][v] = 0
			columns[y][v] = 0
			blocks[n][v] = 0
			board[x][y] = '.'
		}
	}

	dfs(0, 0)
}

func Values(x, y, z [Size]int) (values []int) {
	for i := 0; i < Size; i++ {
		if x[i] == 0 && y[i] == 0 && z[i] == 0 {
			values = append(values, i)
		}
	}
	return values
}

func solveSudoku2(board [][]byte) {
	var (
		rows    = [Size][Size]bool{}
		columns = [Size][Size]bool{}
		blocks  = [Size / Block][Size / Block][Size]bool{}
	)

	for x, row := range board {
		for y, b := range row {
			if b == '.' {
				continue
			}
			v := b - '1'

			rows[x][v] = true
			columns[y][v] = true
			blocks[x/Block][y/Block][v] = true
		}
	}

	var dfs func(x, y int)

	found := false
	dfs = func(x, y int) {
		if y == Size {
			x++
			y = 0
			if x == Size {
				found = true
				return
			}
		}

		if board[x][y] != '.' {
			dfs(x, y+1)
			return
		}

		for i := 0; i < Size; i++ {
			if rows[x][i] || columns[y][i] || blocks[x/Block][y/Block][i] {
				continue
			}

			rows[x][i] = true
			columns[y][i] = true
			blocks[x/Block][y/Block][i] = true
			board[x][y] = byte(i) + '1'

			dfs(x, y+1)
			if found {
				return
			}

			rows[x][i] = false
			columns[y][i] = false
			blocks[x/Block][y/Block][i] = false
			board[x][y] = '.'
		}
	}

	dfs(0, 0)
}

func solveSudoku3(board [][]byte) {
	var (
		rows    = [Size][Size]bool{}
		columns = [Size][Size]bool{}
		blocks  = [Size / Block][Size / Block][Size]bool{}
	)

	type Point struct{ X, Y int }
	var points = make([]Point, 0, Size*Size)
	for x, row := range board {
		for y, b := range row {
			if b == '.' {
				points = append(points, Point{X: x, Y: y})
			} else {
				v := b - '1'
				rows[x][v] = true
				columns[y][v] = true
				blocks[x/Block][y/Block][v] = true
			}
		}
	}

	var dfs func(idx int)

	// 因为题目说明了解唯一, 因此找到一个解后, 直接结束回溯, 实现剪枝
	found := false
	dfs = func(idx int) {
		if idx == len(points) {
			found = true
			return
		}

		x, y := points[idx].X, points[idx].Y
		for i := byte(0); i < Size; i++ {
			if rows[x][i] || columns[y][i] || blocks[x/Block][y/Block][i] {
				continue
			}

			rows[x][i] = true
			columns[y][i] = true
			blocks[x/Block][y/Block][i] = true
			board[x][y] = i + '1'

			dfs(idx + 1)
			if found {
				return
			}

			rows[x][i] = false
			columns[y][i] = false
			blocks[x/Block][y/Block][i] = false
			//board[x][y] = '.' // 会被覆盖, 不用恢复现场
		}
	}

	dfs(0)

	return
}
