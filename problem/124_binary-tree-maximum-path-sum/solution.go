package main

import (
	"math"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
			left  = _max(0, dfs(node.Left))
			right = _max(0, dfs(node.Right))
		)
		ans = _max(ans, node.Val+left+right)
		return node.Val + _max(left, right)
	}
	dfs(root)
	return ans
}
