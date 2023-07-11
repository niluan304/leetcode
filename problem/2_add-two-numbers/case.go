package add_two_numbers

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct{ L1, L2 *ListNode }

type Output *ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				L1: structs.NewListNode([]int{2, 4, 3}),
				L2: structs.NewListNode([]int{5, 6, 4}),
			},
			Except: structs.NewListNode([]int{7, 0, 8}),
		},
		{
			Input: Input{
				L1: structs.NewListNode([]int{0}),
				L2: structs.NewListNode([]int{0}),
			},
			Except: structs.NewListNode([]int{0}),
		},
		{
			Input: Input{
				L1: structs.NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				L2: structs.NewListNode([]int{9, 9, 9, 9}),
			},
			Except: structs.NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
}

type Func func(l1 *ListNode, l2 *ListNode) *ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out Output) {
	return func(in Input) (out Output) {
		return f(in.L1, in.L2)
	}
}

func Equal(l1, l2 Output) bool {
	var (
		n1 *ListNode = l1
		n2 *ListNode = l2
	)

	return tests.Equal(n1.ToSlice(), n2.ToSlice())
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
