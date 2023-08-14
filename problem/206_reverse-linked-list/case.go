package reverse_linked_list

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	head *ListNode
}

type Output *ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head: structs.NewListNode([]int{1, 2, 3, 4, 5}),
			},
			Except: structs.NewListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{1, 2}),
			},
			Except: structs.NewListNode([]int{2, 1}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{12}),
			},
			Except: structs.NewListNode([]int{12}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{}),
			},
			Except: structs.NewListNode([]int{}),
		},
	}
}

type Func func(head *ListNode) *ListNode

func adaptor(f Func) func(in Input) (out Output) {
	return func(in Input) (out Output) {
		root := f(
			in.head,
		)
		return root
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

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
