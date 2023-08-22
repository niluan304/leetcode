package same_tree

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

type Input struct {
	p *structs.TreeNode
	q *structs.TreeNode
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				p: structs.NewTreeByBFS([]any{1, 2, 3}),
				q: structs.NewTreeByBFS([]any{1, 2, 3}),
			},
			Except: true,
		},
		{
			Input: Input{
				p: structs.NewTreeByBFS([]any{1, 2}),
				q: structs.NewTreeByBFS([]any{1, nil, 2}),
			},
			Except: false,
		},
		{
			Input: Input{
				p: structs.NewTreeByBFS([]any{1, 2, 1}),
				q: structs.NewTreeByBFS([]any{1, 1, 2}),
			},
			Except: false,
		},
	}
}

type Func func(p *structs.TreeNode, q *structs.TreeNode) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
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
