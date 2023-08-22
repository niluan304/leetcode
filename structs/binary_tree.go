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

// NewBST 高度平衡二叉树
func NewBST(nodes []int) *TreeNode {
	sort.Ints(nodes)
	return nil
}

func (t *TreeNode) ToSlice() []int {
	return nil
}

// NewTreeByBFS By BFS
func NewTreeByBFS(arr []any) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	rootVal, ok := arr[0].(int)
	if !ok {
		return nil
	}

	root := &TreeNode{
		Val: rootVal,
	}

	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		leftVal, ok := arr[i].(int)
		if ok {
			leftNode := &TreeNode{
				Val: leftVal,
			}
			currNode.Left = leftNode
			queue = append(queue, leftNode)
		}

		i++
		if i >= len(arr) {
			break
		}

		rightVal, ok := arr[i].(int)
		if ok {
			rightNode := &TreeNode{
				Val: rightVal,
			}
			currNode.Right = rightNode
			queue = append(queue, rightNode)
		}

		i++
	}

	return root
}

func (t *TreeNode) BFS() (list []any) {
	var (
		queen = make([]*TreeNode, 0, 100)
	)
	queen = append(queen, t)
	for len(queen) > 0 {
		n := len(queen)

		for _, node := range queen {
			if node == nil {
				list = append(list, nil)
				continue
			}

			list = append(list, node.Val)
			queen = append(queen, node.Left, node.Right)
		}

		queen = queen[n:]
	}

	var n = len(list)
	// 去掉末尾的 nil
	for n > 0 && list[n-1] == nil {
		n--
	}
	return list[:n]
}
