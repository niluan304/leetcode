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

// 题目没说二叉搜索树是平衡的，最坏情况下，会退化为链表，此时单次询问的时间复杂度为 O(n)
// - 时间复杂度：$\mathcal{O}(n \cdot ,m)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// Deprecated: 超时
func closestNodes(root *TreeNode, queries []int) [][]int {
	var n = len(queries)
	var ans = make([][]int, n)
	for i, query := range queries {
		x, y := -1, -1

		floor, ceil := FloorAndCeil(root, query)
		if floor != nil {
			x = floor.Val
		}
		if ceil != nil {
			y = ceil.Val
		}

		ans[i] = []int{x, y}

	}
	return ans
}

// FloorAndCeil
// Floor 小于等于的最大值
// Ceil 大于等于的最小值
func FloorAndCeil(root *TreeNode, key int) (floor, ceil *TreeNode) {
	o := root
	for o != nil {
		switch {
		case key == o.Val:
			return o, o
		case key > o.Val:
			floor = o
			o = o.Right
		case key < o.Val:
			ceil = o
			o = o.Left
		}
	}
	return floor, ceil
}

func closestNodes2(root *TreeNode, queries []int) [][]int {
	nums := make([]int, 0, 1e5)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	n := len(nums)
	ans := make([][]int, len(queries))
	for i, query := range queries {
		mn, mx := -1, -1
		j, ok := slices.BinarySearch(nums, query)
		if j < n {
			mx = nums[j]
		}
		if !ok { // a[j]>q, a[j-1]<q
			j--
		}
		if j >= 0 {
			mn = nums[j]
		}
		ans[i] = []int{mn, mx}
	}
	return ans
}
