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

// 暴力遍历整棵树
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func rangeSumBST(root *TreeNode, low int, high int) int {
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if low <= node.Val && node.Val <= high { // 任意遍历顺序
			ans += node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// bst
// 注：如果你知道线段树，可以看出下面的代码就是线段树的 query。
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func rangeSumBST2(root *TreeNode, low int, high int) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Val > high {
			return dfs(node.Left)
		}
		if node.Val < low {
			return dfs(node.Right)
		}
		// low <= node.Val <= high
		return node.Val + dfs(node.Left) + dfs(node.Right)
	}
	return dfs(root)
}
