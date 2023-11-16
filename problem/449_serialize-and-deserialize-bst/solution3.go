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

type Codec3 struct {
}

func Constructor3() *Codec3 {
	return &Codec3{}
}

const Mod3 = 128

// Serializes a tree to a single string.
// 后序遍历
func (c *Codec3) serialize(root *TreeNode) string {
	var data = make([]byte, 0, 100)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		data = append(data, byte(root.Val/Mod3), byte(root.Val%Mod3))
	}

	dfs(root)
	return string(data)
}

// Deserializes your encoded data to tree.
// 后序遍历, 由于 n 是全局变量, 必须先遍历右子树
func (c *Codec3) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	var idx = len(data)

	//for i := 0; i < len(data); i++ {
	//	fmt.Println(data[i])
	//}

	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if idx == 0 {
			return nil
		}

		val := int(data[idx-2])*Mod3 + int(data[idx-1])
		if val < left || val > right {
			return nil
		}
		idx -= 2
		return &TreeNode{
			Val:   val,
			Right: dfs(val, right), // 后序遍历, 由于 idx 是全局变量, 必须先遍历右子树
			Left:  dfs(left, val),
		}
	}
	return dfs(math.MinInt64, math.MaxInt64)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
