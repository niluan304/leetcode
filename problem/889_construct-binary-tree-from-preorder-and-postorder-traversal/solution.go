package main

import (
	"slices"

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

// todo 题目保证没有重复元素，如果有重复元素，该怎么做？

// dfs
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	var dfs func(preorder []int, postorder []int) *TreeNode
	dfs = func(preorder []int, postorder []int) *TreeNode {
		n := len(preorder)
		if n == 0 {
			return nil
		}

		val := preorder[0]
		if n == 1 {
			return &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}

		i := slices.Index(preorder, postorder[n-2])

		root := &TreeNode{
			Val:   val,
			Left:  constructFromPrePost(preorder[1:i], postorder[:i-1]),
			Right: constructFromPrePost(preorder[i:], postorder[i-1:n-1]),
		}
		return root
	}
	return dfs(preorder, postorder)
}

// dfs
// constructFromPrePost 的优化版本，使用 index 预处理 preorder 值和索引的对应关系，空间换时间。
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func constructFromPrePost2(preorder []int, postorder []int) *TreeNode {
	index := map[int]int{}
	for i, v := range preorder {
		index[v] = i
	}
	var dfs func(preL, preR int, postL, postR int) *TreeNode
	dfs = func(preL, preR int, postL, postR int) *TreeNode {
		if preL == preR {
			return nil
		}
		val := preorder[preL]
		if preR == preL+1 {
			return &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}

		i := index[postorder[postR-2]]

		leftSize := i - (preL + 1) // 左子树的大小
		root := &TreeNode{
			Val:   val,
			Left:  dfs(preL+1, i, postL, postL+leftSize),
			Right: dfs(i, preR, postL+leftSize, postR-1),
		}
		return root
	}

	n := len(preorder)
	return dfs(0, n, 0, n)
}
