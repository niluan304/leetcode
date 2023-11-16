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

// 通过 map 保存节点的 parent 节点
// 获取树最深结点列表
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	type Nodes struct {
		Depth int
		Nodes map[*TreeNode]struct{}
	}

	var parents = map[*TreeNode]*TreeNode{}
	var nodes Nodes
	var dfs func(parent, node *TreeNode, depth int)
	dfs = func(parent, node *TreeNode, depth int) {
		if node == nil {
			// 重置树的最深结点
			if depth > nodes.Depth {
				nodes = Nodes{
					Depth: depth,
					Nodes: map[*TreeNode]struct{}{},
				}
			}
			// 更新树的最深结点
			if depth == nodes.Depth {
				nodes.Nodes[parent] = struct{}{}
			}
			return
		}

		parents[node] = parent
		depth++
		dfs(node, node.Left, depth)
		dfs(node, node.Right, depth)
	}

	dfs(nil, root, 0)

	// 寻找最深结点们的最近公共祖先
	var children = nodes.Nodes
	for len(children) > 1 {
		var p = map[*TreeNode]struct{}{}
		for child := range children {
			p[parents[child]] = struct{}{}
		}
		children = p
	}

	// 返回 "children[0]"
	var ans *TreeNode
	for node := range children {
		ans = node
	}
	return ans
}
