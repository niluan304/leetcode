package merge_two_sorted_lists

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	list1 *ListNode
	list2 *ListNode
}

type Output *ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				list1: structs.NewListNode([]int{1, 2, 4}),
				list2: structs.NewListNode([]int{1, 3, 4}),
			},
			Except: structs.NewListNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			Input: Input{
				list1: structs.NewListNode([]int{}),
				list2: structs.NewListNode([]int{}),
			},
			Except: structs.NewListNode([]int{}),
		},
		{
			Input: Input{
				list1: structs.NewListNode([]int{}),
				list2: structs.NewListNode([]int{0}),
			},
			Except: structs.NewListNode([]int{0}),
		},
	}
}

type Func func(list1 *ListNode, list2 *ListNode) *ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.list1,
			in.list2,
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
	var x1 *structs.ListNode = x
	var y1 *structs.ListNode = y

	return reflect.DeepEqual(x1.ToSlice(), y1.ToSlice())
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
