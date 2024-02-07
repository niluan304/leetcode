package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var ans = 0
	var dfs func(node *TreeNode) (left, right int)

	dfs = func(node *TreeNode) (left, right int) {
		if node == nil {
			return -1, -1 // 表示父节点没有子节点，因为父节点的左右子节点距离为 0
		}

		left = max(dfs(node.Left)) + 1
		right = max(dfs(node.Right)) + 1

		ans = max(ans, left+right)
		return
	}

	dfs(root)
	return ans
}
