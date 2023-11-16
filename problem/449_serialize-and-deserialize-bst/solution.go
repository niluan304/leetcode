package main

import (
	"strconv"
	"strings"

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

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

const (
	Nil = ""
	Sep = ","
)

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var (
		queen = make([]*TreeNode, 0, 100)
		list  = make([]string, 0, 100)
	)
	queen = append(queen, root)
	for len(queen) > 0 {
		n := len(queen)

		for _, node := range queen {
			if node == nil {
				list = append(list, Nil)
				continue
			}

			list = append(list, strconv.Itoa(node.Val))
			queen = append(queen, node.Left, node.Right)
		}

		queen = queen[n:]
	}

	var n = len(list)
	// 去掉末尾的 nil
	for n > 0 && list[n-1] == Nil {
		n--
	}
	return strings.Join(list, Sep)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == Nil {
		return nil
	}

	arr := strings.Split(data, Sep)
	v, _ := strconv.Atoi(arr[0])
	root := &TreeNode{
		Val: v,
	}

	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if arr[i] != Nil {
			val, _ := strconv.Atoi(arr[i])
			leftNode := &TreeNode{
				Val: val,
			}
			currNode.Left = leftNode
			queue = append(queue, leftNode)
		}

		i++
		if i >= len(arr) {
			break
		}

		if arr[i] != Nil {
			val, _ := strconv.Atoi(arr[i])
			rightNode := &TreeNode{
				Val: val,
			}
			currNode.Right = rightNode
			queue = append(queue, rightNode)
		}

		i++
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
