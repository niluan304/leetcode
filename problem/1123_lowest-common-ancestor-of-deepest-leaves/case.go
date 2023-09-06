package lowest_common_ancestor_of_deepest_leaves

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

type Output *structs.TreeNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			},
			Except: structs.NewTreeByBFS([]any{2, 7, 4}),
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{1}),
			},
			Except: structs.NewTreeByBFS([]any{1}),
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{0, 1, 3, nil, 2}),
			},
			Except: structs.NewTreeByBFS([]any{2}),
		},
	}
}

type Func func(root *structs.TreeNode) *structs.TreeNode

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
