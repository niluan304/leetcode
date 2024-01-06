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
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for node, next := head, head.Next; next != nil; node, next = next, next.Next {
		node.Next = &ListNode{
			Val:  gcd(node.Val, next.Val),
			Next: next,
		}
	}
	return head
}

// 求两个数的最大公约数
func gcd(a, b int) int {
	for b != 0 {
		// 使用辗转相除法求最大公约数
		a, b = b%a, a
	}
	return a
}
