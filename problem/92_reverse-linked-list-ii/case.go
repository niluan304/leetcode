package reverse_linked_list_ii

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type Input struct {
	head  []int
	left  int
	right int
}

type Output *ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head:  []int{1, 2, 3, 4, 5},
				left:  2,
				right: 4,
			},
			Except: structs.NewListNode([]int{1, 4, 3, 2, 5}),
		},
		{
			Input: Input{
				head:  []int{5},
				left:  1,
				right: 1,
			},
			Except: structs.NewListNode([]int{5}),
		},
	}
}

type Func func(head *ListNode, left int, right int) *ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			structs.NewListNode(in.head),
			in.left,
			in.right,
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
	x1 := (*structs.ListNode)(x).ToSlice()
	y1 := (*structs.ListNode)(y).ToSlice()

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
