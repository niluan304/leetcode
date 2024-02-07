package main

import (
	"math"

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
func maxPathSum(root *TreeNode) int {
	var ans = math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		var (
			left  = max(0, dfs(node.Left))
			right = max(0, dfs(node.Right))
		)
		ans = max(ans, node.Val+left+right)
		return node.Val + max(left, right)
	}
	dfs(root)
	return ans
}
