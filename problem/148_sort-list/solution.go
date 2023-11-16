package main

import (
	"sort"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var arr []int
	for node := head; node != nil; node = node.Next {
		arr = append(arr, node.Val)
	}

	sort.Ints(arr)

	var i = 0
	for node := head; node != nil; node = node.Next {
		node.Val = arr[i]
		i++
	}

	return head
}

// sortList2 效率低于 sortList
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Next = nodes[i]
	}

	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sortListNode(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sortListNode(head, mid), sortListNode(mid, tail))
}

func leetcode1(head *ListNode) *ListNode {
	return sortListNode(head, nil)
}

func leetcode2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}
