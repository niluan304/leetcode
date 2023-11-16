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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := head
	cur := head.Next

	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	for node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return head
}
