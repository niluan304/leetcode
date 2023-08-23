package binary_tree_right_side_view

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

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{1, 2, 3, nil, 5, nil, 4}),
			},
			Except: []int{1, 3, 4},
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{1, nil, 3}),
			},
			Except: []int{1, 3},
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{}),
			},
			Except: []int{},
		},
	}
}

type Func func(root *structs.TreeNode) []int

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
