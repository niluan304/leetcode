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
func buildTree(inorder []int, postorder []int) *TreeNode {
	var dfs func(inorder []int, postorder []int) *TreeNode
	dfs = func(inorder []int, postorder []int) *TreeNode {
		if len(inorder) == 0 {
			return nil
		}

		m := len(postorder)
		val := postorder[m-1]
		i := slices.Index(inorder, val)

		return &TreeNode{
			Val:   val,
			Left:  dfs(inorder[:i], postorder[:i]),
			Right: dfs(inorder[i+1:], postorder[i:m-1]),
		}
	}

	return dfs(inorder, postorder)
}

// dfs
// buildTree 的优化版本，使用 index 预处理 inorder 值和索引的对应关系，空间换时间。
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func buildTree2(inorder []int, postorder []int) *TreeNode {
	index := map[int]int{}
	for i, v := range inorder {
		index[v] = i
	}

	var dfs func(inL, inR int, postL, postR int) *TreeNode
	dfs = func(inL, inR int, postL, postR int) *TreeNode {
		if inL == inR {
			return nil
		}

		val := postorder[postR-1]
		i := index[val]
		leftSize := i - inL // 左子树的大小
		root := &TreeNode{
			Val:   val,
			Left:  dfs(inL, i, postL, postL+leftSize),
			Right: dfs(i+1, inR, postL+leftSize, postR-1),
		}
		return root
	}

	n := len(inorder)
	return dfs(0, n, 0, n)
}
