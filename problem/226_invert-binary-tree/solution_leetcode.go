package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// leetcode 1 自顶向下遍历
func leetcode1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := leetcode1(root.Left)
	right := leetcode1(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// leetcode 2 前序/后序 + 迭代
func leetcode2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		// 出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root
}

// leetcode 3 前序/后序 + 递归
func leetcode3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	leetcode3(root.Left)
	leetcode3(root.Right)
	return root
}
