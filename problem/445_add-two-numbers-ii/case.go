package add_two_numbers_ii

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	l1 *ListNode
	l2 *ListNode
}

type Output *structs.ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				l1: structs.NewListNode([]int{7, 2, 4, 3}),
				l2: structs.NewListNode([]int{5, 6, 4}),
			},
			Except: structs.NewListNode([]int{7, 8, 0, 7}),
		},
		{
			Input: Input{
				l1: structs.NewListNode([]int{2, 4, 4}),
				l2: structs.NewListNode([]int{5, 6, 3}),
			},
			Except: structs.NewListNode([]int{8, 0, 7}),
		},
		{
			Input: Input{
				l1: structs.NewListNode([]int{0}),
				l2: structs.NewListNode([]int{0}),
			},
			Except: structs.NewListNode([]int{0}),
		},
	}
}

type Func func(l1 *structs.ListNode, l2 *structs.ListNode) *structs.ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.l1,
			in.l2,
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
