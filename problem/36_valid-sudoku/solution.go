package main

const (
	empty = '.'
	size  = 9
)

func isValidSudoku(board [][]byte) bool {
	// check row
	for _, row := range board {
		var list [size]byte
		for _, n := range row {
			if n == empty {
				continue
			}

			idx := n - '1'
			if list[idx] == 0 {
				list[idx]++
			} else {
				return false
			}
		}
	}

	// check column
	for i := 0; i < size; i++ {
		var list [size]byte
		for j := 0; j < size; j++ {
			n := board[j][i]
			if n == empty {
				continue
			}

			idx := n - '1'
			if list[idx] == 0 {
				list[idx]++
			} else {
				return false
			}
		}
	}

	// check block
	for i := 0; i < size; i += 3 {
		for j := 0; j < size; j += 3 {
			var list [size]byte
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					n := board[x][y]
					if n == empty {
						continue
					}

					idx := n - '1'
					if list[idx] == 0 {
						list[idx]++
					} else {
						return false
					}
				}
			}
		}
	}

	return true
}

func isValidSudoku2(board [][]byte) bool {
	type Point struct{ X, Y, A, B int }

	var points [size][]Point
	for x, row := range board {
		for y, num := range row {
			if num == empty {
				continue
			}

			idx := num - '1'

			a, b := x/3, y/3
			for _, point := range points[idx] {
				if point.X == x || // 是否同行
					point.Y == y || // 是否同列
					(point.A == a && point.B == b) { // 是否同块
					return false
				}
			}

			points[idx] = append(points[idx], Point{X: x, Y: y, A: a, B: b})
		}
	}

	return true
}
