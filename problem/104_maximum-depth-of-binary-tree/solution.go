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
func maxDepth(root *TreeNode) int {
	var depth = 0

	var dfs func(node *TreeNode, d int)
	dfs = func(node *TreeNode, d int) {
		if node == nil {
			depth = max(depth, d)
			return
		}

		// 递归
		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}

	dfs(root, 0)

	return depth
}
