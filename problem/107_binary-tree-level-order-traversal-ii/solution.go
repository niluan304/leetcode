package main

import (
	"slices"

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

func levelOrderBottom(root *TreeNode) [][]int {
	var ans = levelOrder(root)
	slices.Reverse(ans)
	return ans
}

// bfs 层序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// LC102. https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int

	queue := []*TreeNode{root}
	for {
		tmp := queue
		queue = nil

		path := make([]int, 0, len(tmp)) // 预分配空间
		for _, node := range tmp {
			if node == nil {
				continue
			}
			path = append(path, node.Val)
			queue = append(queue, node.Left, node.Right)
		}

		if len(path) == 0 {
			break
		}
		ans = append(ans, path)
	}
	return ans
}
