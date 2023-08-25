package insufficient_nodes_in_root_to_leaf_paths

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

type Input struct {
	root  *structs.TreeNode
	limit int
}

type Output *structs.TreeNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				root:  structs.NewTreeByBFS([]any{1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14}),
				limit: 1,
			},
			Except: structs.NewTreeByBFS([]any{1, 2, 3, 4, nil, nil, 7, 8, 9, nil, 14}),
		},
		{
			Input: Input{
				root:  structs.NewTreeByBFS([]any{5, 4, 8, 11, nil, 17, 4, 7, 1, nil, nil, 5, 3}),
				limit: 22,
			},
			Except: structs.NewTreeByBFS([]any{5, 4, 8, 11, nil, 17, 4, 7, nil, nil, nil, 5}),
		},
		{
			Input: Input{
				root:  structs.NewTreeByBFS([]any{1, 2, -3, -5, nil, 4, nil}),
				limit: -1,
			},
			Except: structs.NewTreeByBFS([]any{1, nil, -3, 4}),
		},
	}
}

type Func func(root *structs.TreeNode, limit int) *structs.TreeNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.root,
			in.limit,
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
	x1 := (*TreeNode)(x).BFS()
	y1 := (*TreeNode)(y).BFS()
	return reflect.DeepEqual(x1, y1)
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
