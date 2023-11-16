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
func levelOrder(root *TreeNode) [][]int {
	var (
		ans   = make([][]int, 0, 10)
		queue = []*TreeNode{root}
	)

	for {
		n := len(queue)
		var path = make([]int, 0, n)
		for _, node := range queue {
			if node == nil {
				continue
			}

			path = append(path, node.Val)
			queue = append(queue, node.Left, node.Right)
		}

		queue = queue[n:]
		if len(queue) == 0 {
			break
		}

		ans = append(ans, path)
	}

	return ans
}
