package linked_list_cycle

import (
	"fmt"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

func init() {
	node1.Next.Next.Next.Next = node1.Next
	node2.Next.Next = node2
	node(node1)
	node(node2)
}

func node(node *ListNode) {
	fmt.Println(node.Val)
}

var (
	node1 = &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: nil}}}}
	node2 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	node3 = &ListNode{Val: 1, Next: nil}
)

var cases = func() []tests.Case[*ListNode, bool] {
	return []tests.Case[*ListNode, bool]{
		{Input: node1, Except: true},
		{Input: node2, Except: true},
		{Input: node3, Except: false},
		{Input: nil, Except: false},
	}
}

var funcs = tests.NewFunc[*ListNode, bool]()

func AddCases(c func() []tests.Case[*ListNode, bool]) {
	_cases := cases()
	cases = func() []tests.Case[*ListNode, bool] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func(*ListNode) bool) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

func Bench(b *testing.B) {
	tests.Bench(b, cases, funcs...)
}
