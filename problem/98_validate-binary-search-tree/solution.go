package main

import (
	"math"

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

// dfs 前序遍历
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left int, right int) bool
	dfs = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}

		var v = node.Val
		return (v > left && v < right) && // 判断 node.Val 是否在 (left, right) 之间
			dfs(node.Left, left, v) && // 将左子树的最大值更新为 node.Val
			dfs(node.Right, v, right) // 将右子树的最小值更新为 node.Val
	}

	return dfs(root, math.MinInt, math.MaxInt)
}

// dfs 中序遍历
// 中序遍历二叉搜索树, 其结果是一个递增的数组
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func isValidBST2(root *TreeNode) bool {
	var dfs func(node *TreeNode) bool
	var pre = math.MinInt
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs(node.Left) {
			return false
		}

		if node.Val <= pre {
			return false
		}
		// 更新递增数组的最大值
		pre = node.Val

		if !dfs(node.Right) {
			return false
		}
		return true
	}

	return dfs(root)
}

// dfs 后序遍历
// 后序遍历是自底向上计算子问题的过程。想要学好动态规划的话，请务必掌握这个思想。
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func isValidBST3(root *TreeNode) bool {
	type Range struct{ Min, Max int }

	over := Range{Min: math.MinInt, Max: math.MaxInt}
	empty := Range{Min: math.MaxInt, Max: math.MinInt}

	var dfs func(node *TreeNode) Range
	dfs = func(node *TreeNode) Range {
		if node == nil {
			return empty
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		//fmt.Printf("node: %d, left: %+v, right: %+v\n", node.Val, left, right)
		v := node.Val
		if v <= left.Max || v >= right.Min {
			return over
		}
		return Range{
			Min: min(left.Min, v),
			Max: max(right.Max, v),
		}
	}

	return dfs(root) != over
}
