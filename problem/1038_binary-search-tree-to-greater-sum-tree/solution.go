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
func bstToGst(root *TreeNode) *TreeNode {
	var preSum = []int{0}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		preSum = append(preSum, preSum[len(preSum)-1]+node.Val)
		node.Val = preSum[len(preSum)-1]
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func bstToGst2(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, sum *int)
	dfs = func(node *TreeNode, sum *int) {
		if node == nil {
			return
		}
		dfs(node.Right, sum)
		node.Val += *sum
		*sum = node.Val
		dfs(node.Left, sum)
	}
	dfs(root, new(int))
	return root
}

func bstToGst3(root *TreeNode) *TreeNode {
	var sum = 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}
