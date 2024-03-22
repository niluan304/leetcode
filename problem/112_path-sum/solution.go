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
func hasPathSum(root *TreeNode, targetSum int) bool {
	var ans bool

	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum += node.Val
		if node.Left == nil && node.Right == nil {
			ans = ans || sum == targetSum
			return
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, 0)
	return ans
}
