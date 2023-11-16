package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 哈希表记录结点是否出现过
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func hasCycle(head *ListNode) bool {
	var record = make(map[*ListNode]struct{}, 100)

	for head != nil {
		_, ok := record[head]
		if ok {
			return true
		}

		record[head] = struct{}{}
		head = head.Next
	}

	return false
}

// 快慢指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var fast, slow = head.Next, head

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

// 快慢指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 快指针一定能在慢指针走完第一圈前，追上满指针，因此时间复杂度最大为：O(2n)
// 最差情况：pos为0时，快指针走完第一圈时，慢指针进度为50%，慢指针再走50%，快指针恰好追上慢指针。
func hasCycle3(head *ListNode) bool {
	// 乌龟和兔子同时从起点出发
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		// 兔子追上乌龟（套圈），说明有环
		if fast == slow || fast.Next == slow {
			return true
		}

		slow = slow.Next      // 乌龟走一步
		fast = fast.Next.Next // 兔子走两步
	}

	// 访问到了链表末尾，无环
	return false
}
