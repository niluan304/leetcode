package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**6
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
// 奇数层从左向右遍历，偶数层从右向左遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	var (
		ans   = make([][]int, 0, 10) //
		queue = []*TreeNode{root}    //
		desc  = false                // 奇数层的下一层，即偶数层是从右往左的遍历的
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
		if desc {
			x := len(path)
			for i := 0; i < x/2; i++ {
				path[i], path[x-1-i] = path[x-1-i], path[i]
			}
		}

		ans = append(ans, path)
		desc = !desc
	}

	return ans
}

// bfs 层序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 先正常层序遍历，再反转偶数层
func zigzagLevelOrder2(root *TreeNode) [][]int {
	var (
		ans   = make([][]int, 0, 10) //
		queue = []*TreeNode{root}    //
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

	for i := range ans {
		// 只反转 1/3/5...层
		if i%2 == 0 {
			continue
		}
		x := len(ans[i])
		for j := 0; j < x/2; j++ {
			ans[i][j], ans[i][x-1-j] = ans[i][x-1-j], ans[i][j]
		}
	}

	return ans
}
