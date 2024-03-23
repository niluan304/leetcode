package main

import (
	"math"
)

// 跳格子
//
// - 时间复杂度：O(nm \codt (m+n))。
// - 空间复杂度：O(nm)。
// Deprecated: 超时
func bruteForce(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make([][]*int, m)
	for i := range cache {
		cache[i] = make([]*int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i >= m || j >= n {
			return math.MaxInt32
		}
		if i == m-1 && j == n-1 {
			return 1
		}

		v := &cache[i][j]
		if *v != nil {
			return **v
		}
		defer func() { *v = &res }()

		res = math.MaxInt32
		for k := grid[i][j]; k >= 1; k-- {
			res = min(res, dfs(i, j+k)+1, dfs(i+k, j)+1)
		}
		return res
	}

	ans := dfs(0, 0)
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

// 线段树优化的动态规划
//
// 线段树的下标一般从 1 开始，即 [1, len(nums)]
// 本题改为 [0, len(nums)-1] 更好
//
// - 时间复杂度：O(mnlog(m+n))。
// - 空间复杂度：O(mn)。
func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	colTree, rowTree := make([]*SegmentTree, n), make([]*SegmentTree, m)

	inf := make([]int, max(m, n))
	for i := range inf {
		inf[i] = math.MaxInt32
	}
	for i := range colTree {
		colTree[i] = NewSegmentTreeWithNums(inf[:m])
	}
	for i := range rowTree {
		rowTree[i] = NewSegmentTreeWithNums(inf[:n])
	}

	// 格子数 = 步数 + 1
	rowTree[m-1].Update(n-1, 1)
	colTree[n-1].Update(m-1, 1)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			step := grid[i][j]
			if step == 0 {
				continue
			}

			v := min(
				rowTree[i].Query(j+1, j+step),
				colTree[j].Query(i+1, i+step),
			) + 1

			rowTree[i].Update(j, v)
			colTree[j].Update(i, v)
		}
	}

	ans := colTree[0].Query(0, 0)
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

// 将  minimumVisitedCells 二维数组的线段树映射到一维数组
//
// 线段树的下标一般从 1 开始，即 [1, len(nums)]
// 本题改为 [0, len(nums)-1] 更好
//
// - 时间复杂度：O(mnlog(m+n))。
// - 空间复杂度：O(mn)。
func minimumVisitedCells2(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])

	colTree := NewSegmentTree(0, m*n-1)
	rowTree := NewSegmentTree(0, m*n-1)

	// 格子数 = 步数 + 1
	rowTree.Update(m*n-1, 1)
	colTree.Update(m*n-1, 1)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			step := grid[i][j]
			if step == 0 {
				continue
			}

			x, y := i*n, j*m
			v := min(
				// 将二维数组映射到一位数组，还需要取边界，容易踩坑。
				// 单一的线段树，在 init 方法中设置初始值，搭配 merge 方法就可以避免边界判断。
				rowTree.Query(x+j+1, x+min(j+step, n-1)),
				colTree.Query(y+i+1, y+min(i+step, m-1)),
			) + 1

			rowTree.Update(x+j, v)
			colTree.Update(y+i, v)
		}
	}

	ans := colTree.Query(0, 0)
	if ans >= math.MaxInt32 {
		return -1
	}
	return ans
}

// NewSegmentTreeWithNums
// 根据数组初始化线段树，左右边界为：[1, len(nums)]
// merge 是根据左右儿子节点，确定本节点的值
func NewSegmentTreeWithNums(nums []int) *SegmentTree {
	var dfs func(node *SegmentTree)
	dfs = func(node *SegmentTree) {
		left, right := node.left, node.right
		if left == right {
			node.value = nums[left]
			return
		}

		mid := (left + right) >> 1 // 允许负数范围，需要 /2 下取整
		node.lo = NewSegmentTree(left, mid)
		node.ro = NewSegmentTree(mid+1, right)

		dfs(node.lo)
		dfs(node.ro)
		node.value = node.merge(node.lo.value, node.ro.value)
	}

	// 线段树的下标一般从 1 开始，即 [1, len(nums)]
	// 本题改为 [0, len(nums)-1] 更好
	root := NewSegmentTree(0, len(nums)-1)
	dfs(root)
	return root
}

// NewSegmentTree 初始化线段树
// merge 是根据左右儿子节点，确定本节点的值
func NewSegmentTree(left, right int) *SegmentTree {
	tree := &SegmentTree{
		left:  left,
		right: right,

		lo: nil,
		ro: nil,
	}
	tree.value = tree.init()
	return tree
}

// SegmentTree
// 动态开点线段树·其一·单点修改
type SegmentTree struct {
	left, right int // 该点的区间范围：[left,right]

	value int // 根据题目要求，可能求：区间最值 / 区间和，在 init 里设置

	lo, ro *SegmentTree // 左右儿子节点
}

// Query
// 根据配置的 merge 函数，查询区间 [left, right] 的值
func (t *SegmentTree) Query(left, right int) int {
	if t == nil || left > t.right || right < t.left {
		return t.init()
	}
	if left <= t.left && t.right <= right {
		return t.value
	}

	lo, ro := t.lo.Query(left, right), t.ro.Query(left, right)
	return t.merge(lo, ro)
}

// 也常叫做 push down
// 动态开点
func (t *SegmentTree) spread() {
	mid := (t.left + t.right) >> 1 // 允许负数范围，需要 /2 下取整
	if t.lo == nil {
		t.lo = NewSegmentTree(t.left, mid)
	}
	if t.ro == nil {
		t.ro = NewSegmentTree(mid+1, t.right)
	}
}

// 单点修改，没有延迟操作（只有区间修改需要延迟操作）
func (t *SegmentTree) operator(i int, op func(node *SegmentTree)) {
	if t.left == t.right {
		op(t)
		return
	}

	t.spread() // 动态开点

	mid := (t.left + t.right) >> 1 // 允许负数范围，需要 /2 下取整
	if i <= mid {
		t.lo.operator(i, op)
	} else {
		t.ro.operator(i, op)
	}
	t.value = t.merge(t.lo.value, t.ro.value) // 根据题目要求修改
}

// Set 单点修改，修改指定位置的值
//
// 闭包用于返回修改后的节点值
//
// 闭包的参数是节点的修改前的值，节点不存在时为 NewSegmentTree 里的默认值，闭包的返回值为节点修改后的值
func (t *SegmentTree) Set(i int, do func(old int) int) {
	t.operator(i, func(node *SegmentTree) {
		node.value = do(node.value) // 根据题目要求修改
	})
}

// Update 单点修改，修改指定位置的值
func (t *SegmentTree) Update(i int, value int) {
	t.operator(i, func(node *SegmentTree) {
		node.value = value
	})
}

func (t *SegmentTree) merge(lo int, ro int) int {
	return min(lo, ro)
}

func (t *SegmentTree) init() int {
	return math.MaxInt32
}

// TODO
// 解法一：使用单调栈替代线段树
// 解法二：使用贪心+最小堆解决
