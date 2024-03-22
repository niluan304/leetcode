package main

import (
	"math"
	"runtime/debug"

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return math.MaxInt32
		}

		if node.Left == nil && node.Right == nil {
			return 1
		}

		return min(dfs(node.Right), dfs(node.Left)) + 1
	}
	return dfs(root)
}

func init() {
	debug.SetGCPercent(-1)
}
