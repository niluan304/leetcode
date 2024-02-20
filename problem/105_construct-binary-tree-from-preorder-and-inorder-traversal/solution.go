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
func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(preorder []int, inorder []int) *TreeNode
	dfs = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		v := preorder[0]
		i := slices.Index(inorder, v) // 左子树的大小
		return &TreeNode{
			Val:   v,
			Left:  dfs(preorder[1:i+1], inorder[:i]),
			Right: dfs(preorder[i+1:], inorder[i+1:]),
		}
	}

	return dfs(preorder, inorder)
}

// dfs
// buildTree 的优化版本，使用 index 预处理 inorder 值和索引的对应关系，空间换时间。
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func buildTree2(preorder []int, inorder []int) *TreeNode {
	index := map[int]int{}
	for i, v := range inorder {
		index[v] = i
	}

	var dfs func(preL, preR int, inL, inR int) *TreeNode
	dfs = func(preL, preR int, inL, inR int) *TreeNode {
		if inL == inR {
			return nil
		}

		v := preorder[preL]
		i := index[v]
		mid := preL + i - inL + 1 // i - inL 为左子树的大小

		return &TreeNode{
			Val:   v,
			Left:  dfs(preL+1, mid, inL, i),
			Right: dfs(mid, preR, i+1, inR),
		}
	}

	n := len(preorder)
	return dfs(0, n, 0, n)
}
