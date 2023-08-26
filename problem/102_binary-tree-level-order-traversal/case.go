package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

type Input struct {
	root *structs.TreeNode
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{3, 9, 20, nil, nil, 15, 7}),
			},
			Except: Output{{3}, {9, 20}, {15, 7}},
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{1}),
			},
			Except: Output{{1}},
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{}),
			},
			Except: Output{},
		},
	}
}

type Func func(root *structs.TreeNode) [][]int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.root,
		))
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Equal(x, y Output) bool {
	if len(x) == 0 && len(y) == 0 {
		return true
	}

	return reflect.DeepEqual(x, y)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}