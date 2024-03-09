package main

import (
	"slices"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	. "github.com/niluan304/leetcode/container"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 暴力穷举
// - 时间复杂度：O(n \logn)。
// - 空间复杂度：O(n)。
func mergeKLists(lists []*ListNode) *ListNode {
	nodes := make([]int, 0, 1e4)
	for _, list := range lists {
		for node := list; node != nil; node = node.Next {
			nodes = append(nodes, node.Val)
		}
	}

	slices.Sort(nodes)
	dummy := &ListNode{}
	cur := dummy
	for _, node := range nodes {
		cur.Next = &ListNode{Val: node, Next: nil}
		cur = cur.Next
	}
	return dummy.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	data := make([]*ListNode, 0, len(lists))
	for _, head := range lists {
		if head != nil {
			data = append(data, head)
		}
	}

	hp := NewHeap(data, func(x, y *ListNode) bool {
		return x.Val < y.Val
	})

	dummy := &ListNode{}
	cur := dummy
	for !hp.Empty() {
		node := hp.Pop()      // 剩余节点中的最小节点
		if node.Next != nil { // 下一个节点不为空
			hp.Push(node.Next) // 下一个节点有可能是最小节点，入堆
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next
}
