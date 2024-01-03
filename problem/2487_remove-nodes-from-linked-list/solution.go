package main

import (
	"math"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 将链表转化为列表
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func removeNodes(head *ListNode) *ListNode {
	var stack []int
	for node := head; node != nil; node = node.Next {
		val := node.Val
		m := len(stack)
		for m > 0 && stack[m-1] < val {
			m--
		}
		stack = append(stack[:m], node.Val)
	}

	return NewListNode(stack)
}

func NewListNode(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		cur = cur.Next
	}

	return head.Next
}

// 递归
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func removeNodes2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := removeNodes2(head.Next)
	var val = math.MinInt
	if node != nil {
		val = node.Val
	}

	if head.Val < val {
		return node
	}

	head.Next = node
	return head
}

// 两次反转链表
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeNodes3(head *ListNode) *ListNode {
	root := reverseList(head)
	var pre = root
	for cur := root.Next; cur != nil; cur = cur.Next {
		if pre.Val > cur.Val {
			pre.Next = nil
			continue
		}
		pre.Next = cur
		pre = cur
	}

	return reverseList(root)
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
