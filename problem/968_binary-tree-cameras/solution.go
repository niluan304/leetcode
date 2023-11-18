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
func minCameraCover(root *TreeNode) int {
	type Value struct {
		Self     int
		Father   int
		Children int
	}

	var dfs func(node *TreeNode) Value
	dfs = func(node *TreeNode) Value {
		if node == nil {
			return Value{Self: math.MaxInt32, Father: 0, Children: 0}
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		return Value{
			Self:     min(l.Self, l.Children, l.Father) + min(r.Self, r.Father, r.Children) + 1,
			Father:   min(l.Self, l.Children) + min(r.Self, r.Children),
			Children: min(l.Self+r.Self, l.Self+r.Children, l.Children+r.Self),
		}
	}

	v := dfs(root)
	return min(v.Self, v.Children)
}

func minCameraCover2(root *TreeNode) int {
	type Value struct {
		Self     int
		Father   int
		Children int
	}

	var dfs func(node *TreeNode) Value
	dfs = func(node *TreeNode) Value {
		if node == nil {
			return Value{Self: math.MaxInt32, Father: 0, Children: 0}
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		//self := min(l.Self, l.Children, l.Father) + min(r.Self, r.Father, r.Children) + 1
		//fa := min(l.Self, l.Children) + min(r.Self, r.Children)
		//children := fa + max(0, min(l.Self-l.Children, r.Self-l.Children))

		// 遍历的写法
		var self, fa, minChild = 0, 0, math.MaxInt32
		for _, v := range []Value{l, r} {
			self += min(v.Self, v.Children, v.Father)
			fa += min(v.Self, v.Children)
			minChild = min(minChild, v.Self-v.Children)
		}

		return Value{
			Self:     self + 1, // cost[x] = 1
			Father:   fa,
			Children: fa + max(0, minChild), // 父节点的基础上，将 代价最低为 0
		}
	}

	v := dfs(root)
	return min(v.Self, v.Children)
}
