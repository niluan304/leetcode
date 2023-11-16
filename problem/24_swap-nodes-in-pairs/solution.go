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

func swapPairs(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{Next: head}
		p0    = dummy
		cur   = p0.Next
	)

	const K = 2
	for cur != nil {
		var h = cur
		for i := 0; i < K; i++ {
			if h == nil {
				return dummy.Next
			}
			h = h.Next
		}

		pre := p0.Next
		for i := 0; i < K; i++ {
			next := cur.Next
			cur.Next = pre

			pre = cur
			cur = next
		}

		// p1: 在第i条反转链表中: 反转前的第一个元素, 反转后的最后一个元素,
		// 对于第i+1条反转链表: 上级元素
		p1 := p0.Next
		p1.Next = cur
		p0.Next = pre
		p0 = p1
	}

	return dummy.Next
}
