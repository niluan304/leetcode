package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func reverseList(head *ListNode) *ListNode {
	var arr []int
	for node := head; node != nil; node = node.Next {
		arr = append(arr, node.Val)
	}

	node := head
	for i := len(arr) - 1; i >= 0; i-- {
		node.Val = arr[i]
		node = node.Next
	}
	return head
}

func reverseList2(head *ListNode) *ListNode {
	var root *ListNode
	next := head
	for next != nil {
		node := next
		next = next.Next

		node.Next = root
		root = node
	}

	return root
}

func reverseList3(head *ListNode) *ListNode {
	var (
		root = &ListNode{
			Val:  0,
			Next: nil,
		}
	)

	for head != nil {
		node := head     // 拷贝指针
		head = head.Next // 移动指针

		node.Next = root.Next
		root.Next = node
	}

	return root.Next
}
