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
// 事件复杂度：O(n)
// 空间复杂度：O(n)
func maxAncestorDiff(root *TreeNode) int {
	var ans = 0
	var dfs func(node *TreeNode, minVal int, maxVal int)
	dfs = func(node *TreeNode, minVal int, maxVal int) {
		if node == nil {
			ans = max(ans, maxVal-minVal)
			return
		}

		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)
		dfs(node.Left, minVal, maxVal)
		dfs(node.Right, minVal, maxVal)
	}

	dfs(root, root.Val, root.Val)
	return ans
}
