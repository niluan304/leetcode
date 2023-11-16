package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		lRob, lNotRob := dfs(node.Left)
		rRob, rNotRob := dfs(node.Right)

		curRob := node.Val + lNotRob + rNotRob
		curNotRob := max(lRob, lNotRob) + max(rRob, rNotRob)
		return curRob, curNotRob
	}
	return max(dfs(root))
}

// dfs 后序遍历, 自底向上
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func rob2(root *TreeNode) int {
	var dfs func(node *TreeNode) int

	var cache = map[*TreeNode]int{nil: 0}
	dfs = func(node *TreeNode) int {
		// 已设置 nil: 0, 等价于：
		//if node == nil {
		//	return 0
		//}
		if v, ok := cache[node]; ok {
			return v
		}

		// 递推公式：
		// 选择 node.Val, 不选儿子结点，但还需要加上孙子节点
		// 不选 node.Val, 左右儿子结点相加

		v := node.Val
		if node.Right != nil {
			v += dfs(node.Right.Left) + dfs(node.Right.Right)
		}
		if node.Left != nil {
			v += dfs(node.Left.Left) + dfs(node.Left.Right)
		}
		v = max(v, dfs(node.Right)+dfs(node.Left))
		cache[node] = v
		return v
	}
	return dfs(root)
}
