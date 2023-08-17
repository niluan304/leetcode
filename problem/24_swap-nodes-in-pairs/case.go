package swap_nodes_in_pairs

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	head []int
}

type Output *structs.ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head: []int{1, 2, 3, 4},
			},
			Except: structs.NewListNode([]int{2, 1, 4, 3}),
		},
		{
			Input: Input{
				head: []int{1},
			},
			Except: structs.NewListNode([]int{1}),
		},
		{
			Input: Input{
				head: []int{},
			},
			Except: structs.NewListNode([]int{}),
		},
	}
}

type Func func(head *structs.ListNode) *structs.ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			structs.NewListNode(in.head),
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
