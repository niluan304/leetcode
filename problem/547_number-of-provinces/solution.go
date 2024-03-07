package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	uf := NewUnionFind(n)
	for i, row := range isConnected {
		for j, v := range row {
			if v == 1 {
				uf.Merge(i, j)
			}
		}
	}

	return uf.Groups
}

type UnionFind struct {
	Fa     []int // 代表元素
	Groups int   // 连通分量个数
}

func NewUnionFind(n int) UnionFind {
	fa := make([]int, n) // n+1
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	return UnionFind{Fa: fa, Groups: n}
}

// Find
// 非递归版本
func (u UnionFind) Find(x int) int {
	root := x
	for u.Fa[root] != root {
		root = u.Fa[root]
	}
	for u.Fa[x] != root {
		u.Fa[x], x = root, u.Fa[x]
	}
	return root
}

// FindR
// 递归版本
func (u UnionFind) FindR(x int) int {
	if u.Fa[x] != x {
		u.Fa[x] = u.FindR(u.Fa[x])
	}
	return u.Fa[x]
}

// Merge
// newRoot = -1 表示未发生合并
func (u *UnionFind) Merge(from, to int) (newRoot int) {
	x, y := u.Find(from), u.Find(to)
	if x == y {
		return -1
	}
	u.Fa[x] = y
	u.Groups--
	return y
}

func (u UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
