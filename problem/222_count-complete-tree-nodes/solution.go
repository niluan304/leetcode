package main

import (
	"sort"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	count := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		count++
		dfs(node.Right)
		dfs(node.Left)
	}

	dfs(root)
	return count
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := 0
	for node := root; node.Left != nil; node = node.Left {
		height++
	}

	return sort.Search(1<<(height+1), func(i int) bool {
		if i <= 1<<height { // 用于第一次二分，让边界直接跳到最左边的叶子的节点
			return false
		}
		bits := 1 << (height - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&i == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
