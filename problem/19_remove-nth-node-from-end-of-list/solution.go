package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for next := head; next != nil; next = next.Next {
		length++
	}

	dump := &ListNode{Next: head}
	node := dump
	for i := 0; i < length-n; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next

	return dump.Next
}

// 双指针
// 1. right 先走 n+1 步， 此时 right 和 left 之间相差 n 个节点
// 2. right 和 left 同时走，直到 right 走到尾部，left.Next 就是倒数第 n 个节点
// 3. 删除 left.Next
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var (
		dummy = &ListNode{Next: head} // 哨兵节点, 删除头节点时候，需要用到

		right, left = dummy, dummy // 双指针
	)

	for i := 0; i < n+1; i++ {
		right = right.Next
	}

	// 等价于
	// for i := 0; i < n; i++ {
	//		right = right.Next
	// }
	// for right.Next != nil {
	for right != nil {
		right = right.Next
		left = left.Next
	}

	// 删除倒数第 n 个节点
	left.Next = left.Next.Next

	return dummy.Next
}
