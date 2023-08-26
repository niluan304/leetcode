package lowest_common_ancestor_of_a_binary_tree

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

type Input struct {
	root *structs.TreeNode
	p    *structs.TreeNode
	q    *structs.TreeNode
}

type Output *structs.TreeNode

var cases = func() []tests.Case[Input, Output] {
	case1 := structs.NewTreeByBFS([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	case2 := structs.NewTreeByBFS([]any{1, 2})

	return []tests.Case[Input, Output]{
		{
			Input: Input{
				root: case1,
				p:    case1.Left,
				q:    case1.Right,
			},
			Except: case1,
		},
		{
			Input: Input{
				root: case1,
				p:    case1.Left,
				q:    case1.Left.Right.Right,
			},
			Except: case1.Left,
		},
		{
			Input: Input{
				root: case2,
				p:    case2,
				q:    case2.Left,
			},
			Except: case2,
		},
	}
}

type Func func(root *structs.TreeNode, p *structs.TreeNode, q *structs.TreeNode) *structs.TreeNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.root,
			in.p,
			in.q,
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
	return x == y
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
