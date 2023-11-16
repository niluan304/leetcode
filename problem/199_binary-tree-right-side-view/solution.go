package main

import (
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

// bfs 层序遍历/广度优先搜索
// 右侧视角，看到的就是每一层最右边的那个元素
// 时间复杂度： O(n)
// 空间复杂度： O(n)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		ans   = make([]int, 0, 100)
		queen = make([]*TreeNode, 0, 100)
	)

	queen = append(queen, root)
	for len(queen) > 0 {
		n := len(queen)
		ans = append(ans, queen[n-1].Val) // 上一层的最右边的元素

		for _, node := range queen {
			// 优先添加左节点，使得该层最后一个元素为该层最右边的元素
			if node.Left != nil {
				queen = append(queen, node.Left)
			}

			if node.Right != nil {
				queen = append(queen, node.Right)
			}
		}

		queen = queen[n:] // 更新后 len(queen) == 0, 这一层的下级都是 nil, 即没有下一层
	}
	return ans
}

// dfs 深度优先搜索
// 每次优先遍历右节点而非左节点，实际上保证了 "某一层第一次被访问的节点一定在这一层的最右边"，
// 因此只需要能够判断某一层是不是被第一次访问就行，怎么判断？加一个depth参数就可以
// bilibili@11726753
// 时间复杂度： O(n)
// 空间复杂度： O(n)
func rightSideView2(root *TreeNode) []int {
	var ans = make([]int, 0, 100)

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == len(ans) {
			ans = append(ans, node.Val)
		}

		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)
	return ans
}

// bfs 层序遍历/广度优先搜索
// 右侧视角，看到的就是每一层最右边的那个元素
// 时间复杂度： O(n)
// 空间复杂度： O(n)
func rightSideView3(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		ans   = make([]int, 0, 100)
		queen = make([]*TreeNode, 0, 100)
	)

	queen = append(queen, root)
	for len(queen) > 0 {
		n := len(queen)
		ans = append(ans, queen[0].Val) // 上一层的最右边的元素

		for _, node := range queen {
			// 优先添加右侧的节点, 使得下一次的queen[0]为该层最右边的元素
			if node.Right != nil {
				queen = append(queen, node.Right)
			}

			if node.Left != nil {
				queen = append(queen, node.Left)
			}
		}

		queen = queen[n:] // 更新后 len(queen) == 0, 这一层的下级都是 nil, 即没有下一层
	}
	return ans
}
