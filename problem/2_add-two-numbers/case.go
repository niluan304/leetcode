package add_two_numbers

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct{ L1, L2 *ListNode }

var cases = func() []tests.Case[Input, []int] {
	return []tests.Case[Input, []int]{
		{
			Input: Input{
				L1: structs.NewListNode([]int{2, 4, 3}),
				L2: structs.NewListNode([]int{5, 6, 4}),
			},
			Except: []int{7, 0, 8},
		},
		{
			Input: Input{
				L1: structs.NewListNode([]int{0}),
				L2: structs.NewListNode([]int{0}),
			},
			Except: []int{0},
		},
		{
			Input: Input{
				L1: structs.NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				L2: structs.NewListNode([]int{9, 9, 9, 9}),
			},
			Except: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
}

type Func func(l1 *ListNode, l2 *ListNode) *ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) (out []int) {
	return func(in Input) (out []int) {
		root := f(in.L1, in.L2)
		if !checkResult {
			return nil
		}
		return root.ToSlice()
	}
}

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
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
