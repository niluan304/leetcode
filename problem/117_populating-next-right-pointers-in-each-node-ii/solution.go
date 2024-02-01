package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	var head = root

	setNext := func(cur **Node, node *Node) {
		if node == nil {
			return
		}
		if head == nil {
			head = node
		}
		if *cur != nil {
			(*cur).Next = node
		}
		*cur = node
	}

	for head != nil {
		var cur, next *Node
		for next, head = head, nil; next != nil; next = next.Next {
			setNext(&cur, next.Left)
			setNext(&cur, next.Right)
		}
	}
	return root
}

func connect2(root *Node) *Node {
	var head = root
	for head != nil {
		var cur []*Node
		for head != nil {
			if head.Left != nil {
				cur = append(cur, head.Left)
			}
			if head.Right != nil {
				cur = append(cur, head.Right)
			}
			head = head.Next
		}

		var n = len(cur)
		if n == 0 {
			break
		}
		for i := 0; i < n-1; i++ {
			cur[i].Next = cur[i+1]
		}
		head = cur[0]
	}
	return root
}
