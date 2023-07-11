package remove_nth_node_from_end_of_list

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	*ListNode
	N int
}

var cases = func() []tests.Case[Input, []int] {
	return []tests.Case[Input, []int]{
		{
			Input:  Input{ListNode: structs.NewListNode([]int{1, 2, 3, 4, 5}), N: 2},
			Except: []int{1, 2, 3, 5},
		},
		{
			Input:  Input{ListNode: structs.NewListNode([]int{1}), N: 1},
			Except: []int{},
		},
		{
			Input:  Input{ListNode: structs.NewListNode([]int{1, 2}), N: 1},
			Except: []int{1},
		},
	}
}

type Func func(head *ListNode, n int) *ListNode

func adaptor(f Func) func(in Input) (out []int) {
	return func(in Input) (out []int) {
		root := f(in.ListNode, in.N)
		if !checkResult {
			return nil
		}
		return root.ToSlice()
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[Input, []int]) {
	_cases := cases()
	cases = func() []tests.Case[Input, []int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, []int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, []int]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
