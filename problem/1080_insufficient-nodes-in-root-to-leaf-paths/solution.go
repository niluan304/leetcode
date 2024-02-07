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

// dfs
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		sum += node.Val

		var left, right = node.Left, node.Right
		if left == nil && right == nil {
			return sum
		}

		var maxSum = math.MinInt
		if left != nil {
			total := dfs(left, sum)
			if total < limit {
				node.Left = nil
			}
			maxSum = max(maxSum, total)
		}

		if right != nil {
			total := dfs(right, sum)
			if total < limit {
				node.Right = nil
			}
			maxSum = max(maxSum, total)
		}
		return maxSum
	}

	if dfs(root, 0) < limit {
		return nil
	}
	return root
}

// todo 优化效率
func sufficientSubset2(root *TreeNode, limit int) *TreeNode {
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		sum += node.Val

		if node.Left == nil && node.Right == nil {
			return sum
		}

		var v = math.MinInt
		if node.Left != nil {
			left := dfs(node.Left, sum)
			if left < limit {
				node.Left = nil
			}
			v = max(v, left)
		}
		if node.Right != nil {
			right := dfs(node.Right, sum)
			if right < limit {
				node.Right = nil
			}
			v = max(v, right)
		}
		return v
	}

	sum := dfs(root, 0)
	if sum < limit {
		return nil
	}
	return root
}
