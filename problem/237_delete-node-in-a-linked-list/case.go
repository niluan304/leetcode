package delete_node_in_a_linked_list

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	head []int
	node int
}

type Output *structs.ListNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				head: []int{4, 5, 1, 9},
				node: 5,
			},
			Except: structs.NewListNode([]int{4, 1, 9}),
		},
		{
			Input: Input{
				head: []int{4, 5, 1, 9},
				node: 1,
			},
			Except: structs.NewListNode([]int{4, 5, 9}),
		},
	}
}

type Func func(head *structs.ListNode)

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		head := structs.NewListNode(in.head)
		cur := head
		for {
			if cur.Val == in.node {
				f(cur)
				return head
			}
			cur = cur.Next
		}
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
