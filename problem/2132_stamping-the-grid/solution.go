package main

// ## 前置知识：差分数组
//
// 关于一维差分，请看[【算法小课堂】差分数组](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)
//
// 一维差分的思想可以推广至二维，请点击图片放大查看：
//
// ![LC2132-c.png](https://pic.leetcode.cn/1702439895-HZofag-LC2132-c.png)
//
// ## 前置知识：二维前缀和
//
// 请看[【图解】二维前缀和](https://leetcode.cn/circle/discuss/UUuRex/)
//
// ## 思路
//
// 1. 由于邮票可以互相重叠，贪心地想，能放邮票就放邮票。
// 2. 遍历所有能放邮票的位置去放邮票。注意邮票不能覆盖被占据的格子，也不能出界。
// 3. 放邮票的同时，记录每个空格子被多少张邮票覆盖。如果存在一个空格子没被邮票覆盖，则返回 $\texttt{false}$，否则返回 $\texttt{true}$。
//
// ## 细节
//
// 1. 怎么快速判断一个矩形区域可以放邮票？求出 $\textit{grid}$ 的**二维前缀和**，从而 $\mathcal{O}(1)$ 地求出任意矩形区域的元素和。如果一个矩形区域的元素和等于 $0$，就表示该矩形区域的所有格子都是 $0$。
// 2. 假设用一个二维计数矩阵 $\textit{cnt}$ 记录每个空格子被多少张邮票覆盖，那么放邮票时，就需要把 $\textit{cnt}$ 的一个矩形区域都加一。怎么快速实现？可以用**二维差分矩阵** $d$ 来代替 $\textit{cnt}$。矩形区域都加一的操作，转变成 $\mathcal{O}(1)$ 地对 $d$ 中四个位置的更新操作。
// 3. 最后从二维差分矩阵 $d$ 还原出二维计数矩阵 $\textit{cnt}$。类似对一维差分数组求前缀和得到原数组，我们需要**对二维差分矩阵求二维前缀和**。遍历 $\textit{cnt}$，如果存在一个空格子的计数值为 $0$，就表明该空格子没有被邮票覆盖，返回 $\texttt{false}$，否则返回 $\texttt{true}$。代码实现时，可以直接在 $d$ 数组上**原地**计算出 $\textit{cnt}$。
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/stamping-the-grid/solutions/1199642/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/
func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	// 1. 计算 grid 的二维前缀和
	s := NewMatrixSum(grid)

	m, n := len(grid), len(grid[0])
	diff := make([][]int, m+2)
	for i, _ := range diff {
		diff[i] = make([]int, n+2)
	}

	// 2. 计算二维差分
	// 为方便第 3 步的计算，在 d 数组的最上面和最左边各加了一行（列），所以下标要 +1
	for i2 := stampHeight; i2 <= m; i2++ {
		for j2 := stampWidth; j2 <= n; j2++ {
			i1 := i2 - stampHeight
			j1 := j2 - stampWidth
			if s.query(i1, j1, i2, j2) == 0 {
				diff[i1+1][j1+1]++
				diff[i1+1][j2+1]--
				diff[i2+1][j1+1]--
				diff[i2+1][j2+1]++
			}
		}
	}

	// 3. 还原二维差分矩阵对应的计数矩阵（原地计算）
	for i, row := range grid {
		for j, v := range row {
			diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
			if v == 0 && diff[i+1][j+1] == 0 {
				return false
			}
		}
	}

	return true
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
