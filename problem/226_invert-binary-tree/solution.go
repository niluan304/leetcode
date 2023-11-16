package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	arr := []*TreeNode{root}

	for i := 0; i < len(arr); i++ {
		node := arr[i]
		node.Right, node.Left = node.Left, node.Right

		if node.Left != nil {
			arr = append(arr, node.Left)
		}
		if node.Right != nil {
			arr = append(arr, node.Right)
		}
	}

	return root
}
