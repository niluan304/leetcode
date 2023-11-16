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
func isBalanced(root *TreeNode) bool {
	return Height(root) != -1
}

func Height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := Height(node.Left)
	if left == -1 {
		return -1
	}

	right := Height(node.Right)
	if right == -1 {
		return -1
	}

	if _abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func _abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
