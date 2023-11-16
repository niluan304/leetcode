package main

func exist(board [][]byte, word string) bool {
	sign = make([][]byte, len(board))
	for i, row := range board {
		sign[i] = make([]byte, len(row))
	}

	for x, row := range board {
		for y := range row {
			sign[x][y] = 1
			if dfs(board, word, x, y, "") {
				return true
			}
			sign[x][y] = 0
		}
	}

	return false
}

var rounds = [4]struct{ x, y int }{{x: 0, y: 1}, {x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: -1}}

var sign [][]byte

func dfs(board [][]byte, word string, x, y int, track string) bool {
	track += string([]byte{board[x][y]})

	if track == word {
		return true
	}

	// copy form strings.HasPrefix
	if word[:len(track)] != track {
		track = track[:len(track)-1]
		return false
	}

	m, n := len(board), len(board[0])
	for _, round := range rounds {
		tx, ty := x+round.x, y+round.y
		if tx < 0 || tx >= m || ty < 0 || ty >= n || sign[tx][ty] == 1 {
			continue
		}

		sign[tx][ty] = 1
		if dfs(board, word, tx, ty, track) {
			return true
		}
		sign[tx][ty] = 0
	}

	return false
}

func leetcode(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func leetcode2(board [][]byte, word string) bool {
	row, col, wLen := len(board), len(board[0]), len(word)
	target := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	if row*col < wLen {
		return false
	}
	wMap := make(map[byte]int)
	for i := 0; i < wLen; i++ {
		wMap[word[i]]++
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			wMap[board[i][j]]--
		}
	}
	for _, v := range wMap {
		if v > 0 {
			return false
		}
	}
	flag := make([][]bool, row)
	for i := 0; i < row; i++ {
		flag[i] = make([]bool, col)
	}
	var aa func(i, j, l int) bool
	aa = func(i, j, l int) bool {
		if l == len(word) {
			return true
		}
		f := false
		for k := 0; k < 4; k++ {
			tmpI, tmpJ := i+target[k][0], j+target[k][1]
			if (tmpI >= 0 && tmpI < row) && (tmpJ >= 0 && tmpJ < col) && (board[tmpI][tmpJ] == word[l] && !flag[tmpI][tmpJ]) {
				flag[tmpI][tmpJ] = true
				f = aa(tmpI, tmpJ, l+1)
				flag[tmpI][tmpJ] = false
				if f {
					return true
				}
			}
		}
		return f
	}
	f := false
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				flag[i][j] = true
				f = aa(i, j, 1)
				flag[i][j] = false
				if f {
					return f
				}
			}
		}
	}
	return f
}

func leetcode3(board [][]byte, word string) bool {
	result := false
	row := len(board)
	column := len(board[0])
outer:
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if result = search(board, word, i, j, 0); result {
				break outer
			}
		}
	}
	return result
}

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func search(board [][]byte, word string, i, j, idx int) bool {
	if board[i][j] != word[idx] {
		return false
	}
	if idx == len(word)-1 {
		return true
	}
	keep := board[i][j]
	board[i][j] = '*'
	result := false
	for _, oneDir := range dir {
		newI := i + oneDir[0]
		newJ := j + oneDir[1]
		if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
			if board[newI][newJ] != '*' {
				if search(board, word, newI, newJ, idx+1) {
					result = true
					break
				}
			}
		}
	}
	board[i][j] = keep
	return result
}
