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

// bfs 层序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func findBottomLeftValue(root *TreeNode) int {
	var (
		ans   = 0
		queue = []*TreeNode{root}
	)

	for len(queue) != 0 {
		n := len(queue)
		ans = queue[0].Val
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}

	return ans
}

// bfs 层序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func findBottomLeftValue2(root *TreeNode) int {
	var (
		bottom *TreeNode
		queue  = []*TreeNode{root}
	)

	for len(queue) != 0 {
		n := len(queue)
		for _, node := range queue {
			if node == nil {
				continue
			}
			// 从右向左遍历, 最后的节点就是最底层的最左边节点
			bottom = node
			queue = append(queue, node.Right, node.Left)
		}
		queue = queue[n:]
	}

	return bottom.Val
}
