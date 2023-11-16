package main

import (
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
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans = 0
	var dfs func(node *TreeNode, parent int) int
	dfs = func(node *TreeNode, parent int) int {
		if node == nil {
			return 0
		}

		v := node.Val
		x, y := dfs(node.Right, v), dfs(node.Left, v)
		ans = _max(ans, x+y)
		if v == parent {
			return _max(x, y) + 1
		} else {
			return 0
		}
	}

	dfs(root, root.Val+1)
	return ans
}
