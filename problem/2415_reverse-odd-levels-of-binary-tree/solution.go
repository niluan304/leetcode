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
// bfs
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func reverseOddLevels(root *TreeNode) *TreeNode {
	var isEven = false
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var n = len(queue)
		var nodes = make([]*TreeNode, 0, n*2)
		var values = make([]int, n)
		for i, node := range queue {
			if node == nil { // 完全二叉树
				return root
			}
			nodes = append(nodes, node.Left, node.Right)
			values[i] = node.Val
		}
		if isEven {
			for i, node := range queue {
				node.Val = values[n-1-i]
			}
		}
		isEven = !isEven
		queue = nodes
	}

	return root
}

// bfs
// 时间复杂度: O(n)
// 空间复杂度: O(n)
// #### 方法一：广度优先搜索
//
// **思路与算法**
//
// 我们直接按照层次遍历该二叉树，然后将奇数层中的值进行反转。二叉树按照层次遍历是一个基本的广度优先搜索问题，可以参考「[树的层序遍历](https://oi-wiki.org/graph/tree-basic/#%E6%A0%91%E4%B8%8A-bfs)」。在遍历的同时，对每一层进行标记，如果当前该层为奇数层，则将该层中的节点用数组保存起来，然后将该层所有节点的值进行反转即可。
func reverseOddLevels2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var isEven = false
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var n = len(queue)
		var nodes = make([]*TreeNode, 0, n*2)
		for _, node := range queue {
			if node == nil { // 完全二叉树
				return root
			}
			nodes = append(nodes, node.Left, node.Right)
		}

		if isEven {
			for i := 0; i < n/2; i++ {
				x, y := queue[i], queue[n-1-i]
				x.Val, y.Val = y.Val, x.Val
			}
		}

		isEven = !isEven
		queue = nodes
	}
	return root
}
