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
func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var parent = map[*TreeNode]*TreeNode{}
	var pVal = map[*TreeNode]int{}
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			node := queue[i]
			if node == nil {
				continue
			}
			sum += node.Val
			pVal[parent[node]] += node.Val

			queue = append(queue, node.Left, node.Right)
			parent[node.Left] = node
			parent[node.Right] = node
		}
		for i := 0; i < n; i++ {
			node := queue[i]
			if node == nil {
				continue
			}
			node.Val = sum - pVal[parent[node]]
		}

		queue = queue[n:]
	}
	return root
}
