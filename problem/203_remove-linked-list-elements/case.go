package remove_linked_list_elements

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	head *structs.ListNode
	val  int
}

type Output *structs.ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head: structs.NewListNode([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			Except: structs.NewListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{}),
				val:  1,
			},
			Except: structs.NewListNode([]int{}),
		},
		{
			Input: Input{
				head: structs.NewListNode([]int{7, 7, 7, 7}),
				val:  7,
			},
			Except: structs.NewListNode([]int{}),
		},
	}
}

type Func func(head *structs.ListNode, val int) *structs.ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.head,
			in.val,
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
