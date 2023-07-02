package structs

import (
	"sort"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(nodes []int) *TreeNode {

	return nil
}

// NewBST 高度平衡二叉树
func NewBST(nodes []int) *TreeNode {
	sort.Ints(nodes)
	return nil
}

func (t *TreeNode) ToSlice() []int {
	return nil
}
