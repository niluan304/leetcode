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

// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		dfs func(node *TreeNode)
		ans *TreeNode
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if IsCommonAncestor(node, p, q) {
			ans = node
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

// IsCommonAncestor check whether node is common ancestor of p and q
func IsCommonAncestor(root, p, q *TreeNode) bool {
	var dfs func(node *TreeNode)
	var i = 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		switch node {
		case p, q:
			i++
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return i == 2
}

// dfs 后序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (findP, findQ bool)
	var ans *TreeNode

	dfs = func(node *TreeNode) (findP, findQ bool) {
		if node == nil {
			return false, false
		}

		leftP, leftQ := dfs(node.Left)
		rightP, rightQ := dfs(node.Right)

		findP = leftP || rightP
		findQ = leftQ || rightQ

		switch node {
		case p:
			findP = true
		case q:
			findQ = true
		}
		if ans == nil && findP && findQ {
			ans = node
		}
		return findP, findQ
	}

	dfs(root)
	return ans
}
