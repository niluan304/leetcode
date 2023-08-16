package reverse_nodes_in_k_group

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type Input struct {
	head []int
	k    int
}

type Output *structs.ListNode

type ListNode = structs.ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			Except: structs.NewListNode([]int{2, 1, 4, 3, 5}),
		},
		{
			Input: Input{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			Except: structs.NewListNode([]int{3, 2, 1, 4, 5}),
		},
	}
}

type Func func(head *structs.ListNode, k int) *structs.ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			structs.NewListNode(in.head),
			in.k,
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
