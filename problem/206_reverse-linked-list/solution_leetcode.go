package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func leetcode1(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

// leetcode 2 递归实现
func leetcode2(head *ListNode) *ListNode {
	// 递归
	if head == nil {
		return nil
	}
	reNode := new(ListNode)
	// 第一个re作为尾部标志
	reNode.Val = -5001
	Traversal(head, reNode)
	return reNode
}

func Traversal(node *ListNode, reNode *ListNode) {
	newNode := new(ListNode)
	newNode.Val = node.Val

	if reNode.Val == -5001 {
		newNode.Next = nil
	} else {
		// clone reNode
		reNodeClone := new(ListNode)
		reNodeClone.Val = reNode.Val
		reNodeClone.Next = reNode.Next
		newNode.Next = reNodeClone
	}
	// 和上面的克隆一样，只能值传递
	reNode.Val = newNode.Val
	reNode.Next = newNode.Next
	if node.Next == nil {
		return
	}
	Traversal(node.Next, reNode)
}
