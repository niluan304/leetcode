package main

import (
	"math"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// dfs 前序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 二叉搜索树的特点：左子树的所有节点的值均小于根节点的值，右子树的所有节点的值均大于根节点的值
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var x, y = p.Val, q.Val
	var ans *TreeNode

	var dfs func(node *TreeNode, left, right int)
	dfs = func(node *TreeNode, left, right int) {
		if node == nil ||
			x <= left || y <= left ||
			x >= right || y >= right {
			return
		}

		ans = node
		v := node.Val
		dfs(node.Left, left, v)
		dfs(node.Right, v, right)
	}

	dfs(root, math.MinInt, math.MaxInt)
	return ans
}

// dfs 前序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 二叉搜索树的特点：左子树的所有节点的值均小于根节点的值，右子树的所有节点的值均大于根节点的值
// 找到第一个满足 p.Val <= node.Val <= q.Val 的节点即为最近公共祖先
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var x, y = p.Val, q.Val
	if x > y {
		x, y = y, x
	}

	var ans *TreeNode
	var dfs func(node *TreeNode, left, right int)
	dfs = func(node *TreeNode, left, right int) {
		if node == nil || ans != nil {
			return
		}

		v := node.Val
		if x <= v && v <= y {
			ans = node
			return
		}

		dfs(node.Left, left, v)
		dfs(node.Right, v, right)
	}

	dfs(root, math.MinInt, math.MaxInt)
	return ans
}
