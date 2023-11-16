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

type Codec2 struct {
}

func Constructor2() *Codec2 {
	return &Codec2{}
}

const Mod = 128
const Sep2 = " "

// Serializes a tree to a single string.
// 普通二叉树的序列化我们需要加上空结点才能保证最后反序列化得到的树是唯一的。但是这里是二叉搜索树，它的中序遍历是有序的。有了中序遍历之后，任意给我们前序遍历或者后序遍历，我们都可以构造一个唯一的二叉树。所以二叉搜索树不需要序列化空结点。
//
// 作者：乐乐
// 链接：https://leetcode.cn/problems/serialize-and-deserialize-bst/solutions/1141127/li-yong-liao-er-cha-sou-suo-shu-xing-zhi-2wno/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func (c *Codec2) serialize(root *TreeNode) string {
	var data = make([]byte, 0, 100)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		data = append(data, byte(root.Val/Mod), byte(root.Val%Mod))
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return string(data)
}

// Deserializes your encoded data to tree.
func (c *Codec2) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var idx = 0

	//for i := 0; i < len(data); i++ {
	//	fmt.Println(data[i])
	//}

	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if idx >= len(data)-1 {
			return nil
		}

		val := int(data[idx])*Mod + int(data[idx+1])
		if val < left || val > right {
			return nil
		}
		idx += 2
		root := &TreeNode{Val: val}
		root.Left = dfs(left, val)
		root.Right = dfs(val, right)
		return root
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
