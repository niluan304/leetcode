package structs

import (
	"fmt"
	"time"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nodes []int) *ListNode {
	head := &ListNode{}
	next := head

	for _, node := range nodes {
		next.Next = &ListNode{Val: node}
		next = next.Next
	}

	return head.Next
}

func (root *ListNode) ToSlice() []int {
	var nodes []int

	t := time.Now()
	node := root
	for node != nil {
		nodes = append(nodes, node.Val)
		node = node.Next

		// TODO 确定是否为环形链表
		if time.Since(t) > time.Second {
			panic(fmt.Sprintf("超时：%+v", len(nodes)))
		}
	}
	return nodes
}
