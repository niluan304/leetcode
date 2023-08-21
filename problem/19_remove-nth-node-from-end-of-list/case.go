package remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type Input struct {
	head *structs.ListNode
	n    int
}

type Output *structs.ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head: structs.NewListNode([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			Except: structs.NewListNode([]int{1, 2, 3, 5}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{1}),
				n:    1,
			},
			Except: structs.NewListNode([]int{}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{1, 2}),
				n:    1,
			},
			Except: structs.NewListNode([]int{1}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{1, 2}),
				n:    2,
			},
			Except: structs.NewListNode([]int{2}),
		},
	}
}

type Func func(head *structs.ListNode, n int) *structs.ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.head,
			in.n,
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
