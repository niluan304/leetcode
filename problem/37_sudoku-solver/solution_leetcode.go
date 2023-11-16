package main

import "math/bits"

// 我们可以考虑按照「行优先」的顺序依次枚举每一个空白格中填的数字，
// 通过递归 + 回溯的方法枚举所有可能的填法。
// 当递归到最后一个空白格后，如果仍然没有冲突，说明我们找到了答案；在递归的过程中，
// 如果当前的空白格不能填下任何一个数字，那么就进行回溯。
//
// 由于每个数字在同一行、同一列、同一个九宫格中只会出现一次，
// 因此我们可以使用 line[i]，column[j]，block[x][y]
// 分别表示第 i 行，第 j 列，第 (x,y)个九宫格中填写数字的情况。
// 在下面给出的三种方法中，我们将会介绍两种不同的表示填写数字情况的方法。
//
// 九宫格的范围为 0≤x≤20以及 0≤y≤20。 具体地，第 i 行第 j 列的格子
// 位于第 (⌊i/3⌋,⌊j/3⌋)个九宫格中，其中 ⌊u⌋表示对 u向下取整。
//
// 由于这些方法均以递归 + 回溯为基础，算法运行的时间（以及时间复杂度）
// 很大程度取决于给定的输入数据， 而我们很难找到一个非常精确的渐进紧界。
// 因此这里只给出一个较为宽松的渐进复杂度上界 O(9^(9×9)) ，
// 即最多有 9×9个空白格，每个格子可以填 [1,9]中的任意整数。
//
// 方法一：回溯
// 思路
//
// 最容易想到的方法是用一个数组记录每个数字是否出现。由于我们可以填写的数字范围为 [1,9]，
// 而数组的下标从 0 开始，因此在存储时，我们使用一个长度为 9 的布尔类型的数组，
// 其中 i 个元素的值为 True，当且仅当数字 i+1 出现过。
// 例如我们用 line[2][3]=True 表示数字 4 在第 2 行已经出现过，
// 那么当我们在遍历到第 2 行的空白格时，就不能填入数字 4。
func leetcode1(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

// 方法二：位运算优化
// 思路与算法
//
// 在方法一中，我们使用了长度为 9 的数组表示每个数字是否出现过。
// 我们同样也可以借助位运算，仅使用一个整数表示每个数字是否出现过。
//
// 具体地，数 b 的二进制表示的第 i 位（从低到高，最低位为第 0 位）为 1，当且仅当数字 i+1 已经出现过。
// 例如当 b 的二进制表示为 (011000100)2时，就表示数字 3，7，8 已经出现过。
func leetcode2(board [][]byte) {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int

	flip := func(i, j int, digit byte) {
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				flip(i, j, digit)
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3]) // 0x1ff 即二进制的 9 个 1
		for ; mask > 0; mask &= mask - 1 {                       // 最右侧的 1 置为 0
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
}

// 方法三：枚举优化
// 思路与算法
//
// 我们可以顺着方法二的思路继续优化下去：
//
// 如果一个空白格只有唯一的数可以填入，
// 也就是其对应的 b 值和 b−1 进行按位与运算后得到 0（即 b 中只有一个二进制位为 1）。
// 此时，我们就可以确定这个空白格填入的数，而不用等到递归时再去处理它。
// 这样一来，我们可以不断地对整个数独进行遍历，将可以唯一确定的空白格全部填入对应的数。
// 随后我们再使用与方法二相同的方法对剩余无法唯一确定的空白格进行递归 + 回溯
func leetcode3(board [][]byte) {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int

	flip := func(i, j int, digit byte) {
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}

	for i, row := range board {
		for j, b := range row {
			if b != '.' {
				digit := b - '1'
				flip(i, j, digit)
			}
		}
	}

	for {
		modified := false
		for i, row := range board {
			for j, b := range row {
				if b != '.' {
					continue
				}
				mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3])
				if mask&(mask-1) == 0 { // mask 的二进制表示仅有一个 1
					digit := byte(bits.TrailingZeros(mask))
					flip(i, j, digit)
					board[i][j] = digit + '1'
					modified = true
				}
			}
		}
		if !modified {
			break
		}
	}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3]) // 0x1ff 即二进制的 9 个 1
		for ; mask > 0; mask &= mask - 1 {                       // 最右侧的 1 置为 0
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
}

// 0 ms 的代码示例
func leetcodeMinTime(board [][]byte) {
	var rm, cm [9][9]int
	var blk [3][3][9]int
	for i, row := range board {
		for j, e := range row {
			if e == '.' {
				continue
			}
			idx := e - '1'
			rm[i][idx]++
			cm[j][idx]++
			blk[i/3][j/3][idx]++
		}
	}
	helper37(board, rm, cm, blk, 0, 0)
}

func helper37(board [][]byte, rm, cm [9][9]int, blk [3][3][9]int, r, c int) bool {
	for ; r < 9; r++ {
		for ; c < 9; c++ {
			if board[r][c] == '.' {
				goto find
			}
		}
		c = 0
	}
	return true
find:
	for i := 0; i < 9; i++ {
		if rm[r][i] == 0 && cm[c][i] == 0 && blk[r/3][c/3][i] == 0 {
			board[r][c] = '1' + byte(i)
			rm[r][i] = 1
			cm[c][i] = 1
			blk[r/3][c/3][i] = 1
			nextr, nextc := r, c
			if nextc == 8 {
				nextr++
				nextc = 0
			} else {
				nextc++
			}
			if helper37(board, rm, cm, blk, nextr, nextc) {
				return true
			}
			board[r][c] = '.'
			rm[r][i] = 0
			cm[c][i] = 0
			blk[r/3][c/3][i] = 0
		}
	}
	return false
}

// 执行用时为 0 ms 的范例
func leetcodeMinTime2(board [][]byte) {
	cols, rows, boxs := [9][10]bool{}, [9][10]bool{}, [9][10]bool{}
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			z := (i/3)*3 + j/3
			v := int(c - '0')
			cols[i][v], rows[j][v], boxs[z][v] = true, true, true
		}
	}
	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		if cur == 81 {
			return true
		}
		x, y := cur/9, cur%9
		if board[x][y] != '.' {
			return dfs(cur + 1)
		}
		z := (x/3)*3 + y/3
		for i := 1; i <= 9; i++ {
			if cols[x][i] || rows[y][i] || boxs[z][i] {
				continue
			}
			cols[x][i], rows[y][i], boxs[z][i] = true, true, true
			board[x][y] = byte('0' + i)
			if dfs(cur + 1) {
				return true
			}
			cols[x][i], rows[y][i], boxs[z][i] = false, false, false
			board[x][y] = '.'
		}
		return false
	}
	dfs(0)
}

// 1.81 MB 的代码示例
func leetcodeMinMemory(board [][]byte) {
	var s = &sudoku{}
	s.Init(board)
	board = s.FillBoard(board)
}

type cell struct {
	i    int
	j    int
	try  byte
	fill byte
}

const (
	Init  byte = 0xff
	Begin byte = 0
	End   byte = 8

	Empty byte = '.'
	Base  byte = '1'
)

type sudoku struct {
	row    [9]int
	col    [9]int
	grid   [9]int
	empty  []cell
	n      int
	cursor int
}

func (s *sudoku) Init(board [][]byte) {
	if s.empty != nil {
		s.empty = s.empty[:0]
	}
	for i := 0; i < 9; i++ {
		s.row[i], s.col[i], s.grid[i] = 0, 0, 0
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == Empty {
				s.empty = append(s.empty, cell{
					i: i, j: j,
					try: Init, fill: Init,
				})
				continue
			}
			s.change(i, j, board[i][j]-Base)
		}
	}
	s.n, s.cursor = len(s.empty), 0
}

// call once is fill (i,j) with byte(Base+c);
// call twice is cancel fill (i,j) with byte (Base+c)
func (s *sudoku) change(i, j int, c byte) {
	s.row[i] ^= 1 << c
	s.col[j] ^= 1 << c
	s.grid[i/3*3+j/3] ^= 1 << c
}

func (s *sudoku) Done() bool {
	return s.cursor == s.n
}

// 返回是否成功填入一个当前有效的值
func (s *sudoku) FillOne() bool {
	// cancel pre filled
	now := &s.empty[s.cursor]

	// cancel fill now
	if now.fill != Init && now.fill == now.try {
		s.change(now.i, now.j, now.fill)
	}

	for now.try++; now.try <= End; now.try++ {
		k := 1 << now.try
		if ((s.row[now.i] & k) | (s.col[now.j] & k) | (s.grid[now.i/3*3+now.j/3] & k)) == 0 {
			now.fill = now.try
			s.change(now.i, now.j, now.fill)
			s.cursor++
			return true
		}
	}

	return false
}

// 无法FillOne的时候，回溯
func (s *sudoku) Backtrack() {
	for s.cursor >= 0 && s.empty[s.cursor].try > End {
		s.empty[s.cursor].try = Init
		s.cursor--
	}
	if s.cursor < 0 {
		panic("error")
	}
}

func (s *sudoku) FillBoard(board [][]byte) [][]byte {
	for !s.Done() {
		if !s.FillOne() {
			s.Backtrack()
		}
	}
	for _, now := range s.empty {
		board[now.i][now.j] = now.fill + Base
	}
	return board
}
