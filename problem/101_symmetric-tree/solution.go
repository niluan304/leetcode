package main

import (
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

// dfs
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isSymmetric(root *TreeNode) bool {
	var dfs func(left, right *TreeNode) bool
	dfs = func(x, y *TreeNode) bool {
		if x == nil && y == nil {
			return true
		}

		if x != nil && y != nil && x.Val == y.Val {
			return dfs(x.Left, y.Right) && dfs(x.Right, y.Left)
		}
		return false
	}

	return dfs(root, root)
}

// dfs
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isSymmetric2(root *TreeNode) bool {
	return DFS(root, root)
}

func DFS(x, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}

	if x != nil && y != nil && x.Val == y.Val {
		return DFS(x.Left, y.Right) && DFS(x.Right, y.Left)
	}
	return false
}
