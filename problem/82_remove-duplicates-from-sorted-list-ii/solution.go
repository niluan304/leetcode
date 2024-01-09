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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cur = head
	var count = map[int]int{}
	for cur != nil {
		count[cur.Val]++
		cur = cur.Next
	}

	cur = head
	var node = &ListNode{}
	list := node
	for cur != nil {
		if count[cur.Val] == 1 {
			node.Next = &ListNode{Val: cur.Val}
			node = node.Next
		}
		cur = cur.Next
	}

	return list.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	var list = &ListNode{}
	var count = map[int]int{}
	var cur = list

	for head != nil {
		count[head.Val]++
		if count[head.Val] == 1 {
			if cur.Next != nil {
				cur = cur.Next
			}

			cur.Next = &ListNode{Val: head.Val}
		} else {
			cur.Next = nil
		}

		head = head.Next
	}

	return list.Next
}

// 1. 找到最后一个相同的节点
// 2. 删除所有相同的节点
// 3. 重复 1，2 步骤
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func deleteDuplicates3(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{Next: head}
		pre   = dummy
		cur   = pre.Next
	)

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			// 找到最后一个相同的节点
			v := cur.Val
			for cur != nil && cur.Val == v {
				cur = cur.Next
			}
			// 删除所有相同的节点
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}

func deleteDuplicates4(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if cur.Next.Next.Val == val {
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
