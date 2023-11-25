package main

import (
	"cmp"

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
func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans = 0
	var list = make([]int, 10)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			list[node.Val] ^= 1
			dfs(node.Left)
			list[node.Val] ^= 1
		}

		if node.Right != nil {
			list[node.Val] ^= 1
			dfs(node.Right)
			list[node.Val] ^= 1
		}

		if node.Right == nil && node.Left == nil {
			list[node.Val] ^= 1
			if Sum(list) <= 1 {
				ans++
			}
			list[node.Val] ^= 1
		}
	}
	dfs(root)
	return ans
}

func pseudoPalindromicPaths2(root *TreeNode) int {
	var list = make([]int, 10)
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		list[node.Val] ^= 1
		defer func() { list[node.Val] ^= 1 }()

		if node.Right == nil && node.Left == nil {
			if Sum(list) <= 1 {
				return 1
			}
			return 0
		}
		return dfs(node.Left) + dfs(node.Right)
	}

	return dfs(root)
}

func pseudoPalindromicPaths3(root *TreeNode) int {
	var ans = 0
	var dfs func(node *TreeNode, mask int)
	dfs = func(node *TreeNode, mask int) {
		if node == nil {
			return
		}

		mask ^= 1 << node.Val
		if node.Left == nil && node.Right == nil { // 等价于 node.Right == node.Left
			if mask&(mask-1) == 0 {
				ans++
			}
			return
		}
		dfs(node.Left, mask)
		dfs(node.Right, mask)
		return
	}

	dfs(root, 0)
	return ans
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	m := x[0]
	for i := 1; i < len(x); i++ {
		m += x[i]
	}

	return m
}
