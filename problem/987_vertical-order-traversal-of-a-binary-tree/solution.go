package main

import (
	"slices"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	. "github.com/niluan304/leetcode/copypasta"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// dfs
//
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func verticalTraversal(root *TreeNode) [][]int {
	type Item struct{ Row, Col, Val int }

	m := map[int][]Item{}
	var dfs func(node *TreeNode, row, col int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}

		m[col] = append(m[col], Item{
			Row: row,
			Col: col,
			Val: node.Val,
		})

		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	keys := MapKeys(m)
	slices.Sort(keys[:])

	ans := make([][]int, len(keys))
	for i, key := range keys {
		slices.SortFunc(m[key][:], func(a, b Item) int {
			if a.Row != b.Row {
				return a.Row - b.Row
			}
			if a.Col != b.Col {
				return a.Col - b.Col
			}
			return a.Val - b.Val
		})

		for _, item := range m[key] {
			ans[i] = append(ans[i], item.Val)
		}
	}

	return ans
}

// dfs
//
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func verticalTraversal2(root *TreeNode) [][]int {
	minCol, maxCol := 0, 0
	maxRow := map[int]int{}

	type pair struct{ row, col int }
	m := map[pair][]int{}
	var dfs func(node *TreeNode, row, col int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}

		minCol = min(minCol, col)
		maxCol = max(maxCol, col)
		maxRow[col] = max(maxRow[col], row)

		key := pair{row: row, col: col}
		m[key] = append(m[key], node.Val) // 按照 [col][row]分组

		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	ans := make([][]int, maxCol-minCol+1)
	for col := minCol; col <= maxCol; col++ {
		for row := 0; row <= maxRow[col]; row++ {
			values := m[pair{row: row, col: col}]
			slices.Sort(values)
			ans[col-minCol] = append(ans[col-minCol], values...)
		}
	}
	return ans
}
