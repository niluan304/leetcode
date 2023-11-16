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
// 快指针一定能在慢指针走完第一圈前，追上满指针，因此时间复杂度最大为：O(2n)
// 最差情况：pos为0时，快指针走完第一圈时，慢指针进度为50%，慢指针再走50%，快指针恰好追上慢指针。
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			cycle := head
			for cycle != slow {
				cycle = cycle.Next
				slow = slow.Next
			}
			return cycle
		}
	}
	return nil
}

// 哈希表保存链表结点的指针
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func detectCycle2(head *ListNode) *ListNode {
	var m = make(map[*ListNode]struct{}, 100)

	for head != nil {
		// 如果存在, 那么就是环的第一个节点
		if _, ok := m[head]; ok {
			return head
		}

		m[head] = struct{}{}
		head = head.Next
	}

	// 访问到了链表末尾，无环
	return nil
}
