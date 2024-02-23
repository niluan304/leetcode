package main

import (
	"slices"
	"sort"

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

// bfs + 排序
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var nums []int

	queue := []*TreeNode{root}
	for {
		num := 0
		tmp := queue
		queue = nil
		for _, node := range tmp {
			if node == nil {
				continue
			}
			num += node.Val
			queue = append(queue, node.Left, node.Right)
		}
		if len(queue) == 0 {
			break
		}
		nums = append(nums, num)
	}

	n := len(nums)
	if n < k {
		return -1
	}
	// 降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return int64(nums[k-1])
}

// dfs + 排序
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func kthLargestLevelSum2(root *TreeNode, k int) int64 {
	var nums []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		_nums := make([]int, max(depth-len(nums)+1, 0))
		nums = append(nums, _nums...)
		nums[depth] += node.Val

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	n := len(nums)
	if n < k {
		return -1
	}
	// 升序排列
	slices.Sort(nums)
	return int64(nums[n-k])
}
