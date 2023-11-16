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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}

	return false
}
