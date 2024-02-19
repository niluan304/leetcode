package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     left *TreeNode
 *     right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	xDepth, yDepth := 0, 0
	xParent, yParent := root, root

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xDepth = depth
			xParent = parent
		}
		if node.Val == y {
			yDepth = depth
			yParent = parent
		}

		dfs(node.Left, node, depth+1)
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xParent != yParent
}

func isCousins2(root *TreeNode, x int, y int) bool {
	queue := []*TreeNode{root}

	parent := map[*TreeNode]*TreeNode{}
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		var xNode, yNode *TreeNode
		for _, node := range tmp {
			if node == nil {
				continue
			}
			if node.Val == x {
				xNode = node
			}
			if node.Val == y {
				yNode = node
			}
			parent[node.Left] = node
			parent[node.Right] = node
			tmp = append(tmp, node.Left, node.Right)
		}

		if xNode == nil && yNode == nil {
			queue = tmp
		} else if xNode != nil && yNode != nil {
			return parent[xNode] != parent[yNode]
		} else {
			return false
		}
	}

	return false
}
