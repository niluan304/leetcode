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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var list = make([]*ListNode, 0, 100)

	for node := head; node != nil; node = node.Next {
		list = append(list, node)
	}

	mid := len(list) / 2
	for i := 0; i < mid; i++ {
		tail := list[len(list)-1-i]

		list[i].Next = tail
		tail.Next = list[i+1]
	}

	list[mid].Next = nil
}
