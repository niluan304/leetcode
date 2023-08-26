package find_bottom_left_tree_value

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

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{2, 1, 3}),
			},
			Except: 1,
		},
		{
			Input: Input{
				root: structs.NewTreeByBFS([]any{1, 2, 3, 4, nil, 5, 6, nil, nil, 7}),
			},
			Except: 7,
		},
	}
}

type Func func(root *structs.TreeNode) int

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
