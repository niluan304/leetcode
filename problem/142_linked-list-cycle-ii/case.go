package linked_list_cycle_ii

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Input struct {
	head *structs.ListNode
	pos  int
}

type Output *structs.ListNode

func NewCase(nums []int, pos int) tests.Case[Input, Output] {
	var head, cycle = structs.NewCycleListNode(nums, pos)
	return tests.Case[Input, Output]{
		Input: Input{
			head: head,
			pos:  pos,
		},
		Except: cycle,
	}
}

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		NewCase([]int{3, 2, 0, -4}, 1),
		NewCase([]int{1, 2}, 0),
		NewCase([]int{1}, -1),
	}
}

type Func func(head *structs.ListNode) *structs.ListNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.head,
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

// Equal 比较指针, 确定是否为同一个对象
func Equal(x, y Output) bool {

	return x == y
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
