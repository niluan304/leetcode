package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法一：模拟
// 思路与算法
//
// 由于输入的两个链表都是逆序存储数字的位数的，因此两个链表中同一位置的数字可以直接相加。
//
// 我们同时遍历两个链表，逐位计算它们的和，并与当前位置的进位值相加。
// 具体而言，如果当前两个链表处相应位置的数字为 n1,n2，进位值为 carry，则它们的和为 n1+n2+carry；
// 其中，答案链表处相应位置的数字为 (n1+n2+carry)mod10，而新的进位值为 ⌊n1+n2+carry⌋/10。
//
// 如果两个链表的长度不同，则可以认为长度短的链表的后面有若干个 0 。
//
// 此外，如果链表遍历结束后，有 carry>0，还需要在答案链表的后面附加一个节点，节点的值为 carry。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/add-two-numbers/solutions/435246/liang-shu-xiang-jia-by-leetcode-solution/
func leetcode1(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
