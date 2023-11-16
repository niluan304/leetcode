package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	next := head

	for list1 != nil && list2 != nil {
		node := 0

		if list1.Val < list2.Val {
			node = list1.Val
			list1 = list1.Next
		} else {
			node = list2.Val
			list2 = list2.Next
		}

		next.Next = &ListNode{Val: node}
		next = next.Next
	}

	list := list1
	if list == nil {
		list = list2
	}
	next.Next = list

	return head.Next
}
