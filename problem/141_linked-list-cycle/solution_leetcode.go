package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func leetcode(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func leetcode2(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	for fast != nil && fast.Next != nil {
		if head == fast {
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}
	return false
}
