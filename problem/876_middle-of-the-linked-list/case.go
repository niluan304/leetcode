package middle_of_the_linked_list

import (
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type Input struct {
	head *ListNode
}

type ListNode = structs.ListNode

type Output *ListNode

var cases = func() []tests.Case[Input, Output] {
	var NewCase = func(list []int) tests.Case[Input, Output] {
		var n = len(list)
		head := structs.NewListNode(list)
		mid := head
		for i := 0; i < n/2; i++ {
			mid = mid.Next
		}
		return tests.Case[Input, Output]{
			Input: Input{
				head: head,
			},
			Except: mid,
		}
	}

	return []tests.Case[Input, Output]{
		NewCase([]int{1, 2, 3, 4, 5}),
		NewCase([]int{1, 2, 3, 4, 5, 6}),
	}
}

type Func func(head *ListNode) *ListNode

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
