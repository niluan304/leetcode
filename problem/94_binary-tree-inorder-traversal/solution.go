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

// dfs 中序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func inorderTraversal(root *TreeNode) []int {
	var ans = make([]int, 0, 100)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}
