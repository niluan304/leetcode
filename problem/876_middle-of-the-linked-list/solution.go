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

// 快慢指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func middleNode(head *ListNode) *ListNode {
	var (
		slow = head
		fast = head
	)

	// fast.Next != nil 用于链表长度为奇数时，slow指向中间节点
	// fast != nil 用于链表长度为偶数时，slow指向中间节点的后一个节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 用数组保存链表
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func middleNode2(head *ListNode) *ListNode {
	var list = make([]*ListNode, 0, 100)

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	return list[len(list)/2]
}
