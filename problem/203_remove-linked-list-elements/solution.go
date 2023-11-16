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

// 迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElements(head *ListNode, val int) *ListNode {
	var (
		dummy = &ListNode{Next: head} // 哨兵结点, 用于删除头结点
		pre   = dummy                 // cur的前驱结点
		cur   = pre.Next              //
	)

	for cur != nil {
		if cur.Val == val {
			// 删除值为 val 的节点
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	return dummy.Next
}

// 迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElements2(head *ListNode, val int) *ListNode {
	var (
		dummy = &ListNode{Next: head}
		cur   = dummy
	)
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
