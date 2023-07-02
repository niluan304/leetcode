package structs

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

	node := root
	for node != nil {
		nodes = append(nodes, node.Val)
		node = node.Next
	}
	return nodes
}
