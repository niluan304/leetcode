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

// 反转链表
// 时间复杂度： O(N)
// 空间复杂度： O(N)
func reverseKGroup(head *ListNode, k int) *ListNode {
	var n = 0
	h := head
	for h != nil {
		n++
		h = h.Next
	}

	n = n / k * k
	var list = make([]*ListNode, n)
	for i := 0; i < n; i++ {
		y := (i/k + 1) * k
		x := y - i%k - 1
		list[x] = head

		head = head.Next
	}

	for i := 1; i < n; i++ {
		list[i-1].Next = list[i]
	}
	list[n-1].Next = head

	return list[0]
}

// 反转链表
// 时间复杂度： O(N)
// 空间复杂度： O(N)
func reverseKGroup2(head *ListNode, k int) *ListNode {
	var list = make([]*ListNode, 0, 100)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	var n = len(list)
	for i := 1; i <= n/k; i++ {
		for j := 0; j < k/2; j++ {
			x1 := i*k - j - 1
			x2 := i*k - (k - j)

			list[x1], list[x2] = list[x2], list[x1]
		}
	}

	for i := 1; i < n; i++ {
		list[i-1].Next = list[i]
	}
	list[n-1].Next = nil
	return list[0]
}

// 反转链表
// 时间复杂度： O(N)
// 空间复杂度： O(1)
func reverseKGroup3(head *ListNode, k int) *ListNode {
	var n = 0
	h := head
	for h != nil {
		n++
		h = h.Next
	}

	var (
		dummy = &ListNode{Next: head}
		p0    = dummy
		cur   = p0.Next
	)

	for i := 0; i < n/k; i++ {
		pre := p0.Next
		for j := 0; j < k; j++ {
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

// 反转链表
// 时间复杂度： O(N)
// 空间复杂度： O(1)
func reverseKGroup4(head *ListNode, k int) *ListNode {
	var (
		dummy = &ListNode{Next: head}
		p0    = dummy
		cur   = p0.Next
	)

	for cur != nil {
		var h = cur
		for i := 0; i < k; i++ {
			if h == nil {
				return dummy.Next
			}
			h = h.Next
		}

		pre := p0.Next
		for i := 0; i < k; i++ {
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
