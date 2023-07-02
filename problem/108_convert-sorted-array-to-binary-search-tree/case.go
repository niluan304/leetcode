package convert_sorted_array_to_binary_search_tree

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

// todo 补全测试用例
var cases = func() []tests.Case[[]int, *TreeNode] {
	return []tests.Case[[]int, *TreeNode]{
		{Input: nil, Except: nil},
	}
}

type Func func([]int) *TreeNode

func adaptor(f Func) func(in []int) (out *TreeNode) {
	return f
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[[]int, *TreeNode]) {
	_cases := cases()
	cases = func() []tests.Case[[]int, *TreeNode] {
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
