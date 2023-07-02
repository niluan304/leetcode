package reverse_linked_list

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

var cases = func() []tests.Case[*ListNode, []int] {
	return []tests.Case[*ListNode, []int]{
		{Input: structs.NewListNode([]int{1, 2, 3, 4, 5}), Except: []int{5, 4, 3, 2, 1}},
		{Input: structs.NewListNode([]int{1, 2}), Except: []int{2, 1}},
		{Input: structs.NewListNode([]int{12}), Except: []int{12}},
		{Input: structs.NewListNode([]int{}), Except: []int{}},
	}
}

type Func func(head *ListNode) *ListNode

func adaptor(f Func) func(in *ListNode) (out []int) {
	return func(in *ListNode) (out []int) {
		root := f(in)
		if !checkResult {
			return nil
		}
		return root.ToSlice()
	}
}

var funcs = tests.NewFuncWithAdaptor(adaptor)

func AddCases(c func() []tests.Case[*ListNode, []int]) {
	_cases := cases()
	cases = func() []tests.Case[*ListNode, []int] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
