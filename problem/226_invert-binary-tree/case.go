package invert_binary_tree

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

var cases = func() []tests.Case[*TreeNode, []int] {
	return []tests.Case[*TreeNode, []int]{
		{Input: NewTreeNode([]int{4, 2, 7, 1, 3, 6, 9}), Except: []int{4, 7, 2, 9, 6, 3, 1}},
		{Input: NewTreeNode([]int{2, 1, 3}), Except: []int{2, 3, 1}},
		{Input: NewTreeNode([]int{}), Except: []int{}},
	}
}

func NewTreeNode(slice []int) *TreeNode {
	switch len(slice) {
	case 3:
		return &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}

	case 7:
		return &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
		}
	}

	return nil
}

type Func func(root *TreeNode) *TreeNode

func adaptor(f Func) func(in *TreeNode) (out []int) {
	return func(in *TreeNode) (out []int) {
		root := f(in)
		if !checkResult {
			return nil
		}
		return result(root)
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func result(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := []*TreeNode{root}
	var ret []int
	for i := 0; i < len(arr); i++ {
		node := arr[i]
		if node.Left != nil {
			arr = append(arr, node.Left)
		}

		if node.Right != nil {
			arr = append(arr, node.Right)
		}

		ret = append(ret, node.Val)
	}
	return ret
}

func AddCases(c func() []tests.Case[*TreeNode, []int]) {
	_cases := cases()
	cases = func() []tests.Case[*TreeNode, []int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
