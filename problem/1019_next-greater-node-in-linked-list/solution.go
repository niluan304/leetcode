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

// 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func nextLargerNodes(head *ListNode) []int {
	var list []int
	for node := head; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	var ans = make([]int, len(list))
	var stack []int
	for i, v := range list {
		for m := len(stack); m > 0; m-- {
			j := stack[m-1]
			if v <= list[j] {
				break
			}
			stack = stack[:m-1]
			ans[j] = v
		}
		stack = append(stack, i)
	}
	return ans
}

// 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func nextLargerNodes2(head *ListNode) []int {
	type Value struct {
		index int
		value int
	}

	var ans []int
	var stack []Value

	for node, i := head, 0; node != nil; node, i = node.Next, i+1 {
		v := node.Val
		for m := len(stack); m > 0; m-- {
			top := stack[m-1]
			if v <= top.value {
				break
			}
			stack = stack[:m-1]
			ans[top.index] = v
		}
		stack = append(stack, Value{index: i, value: v})
		ans = append(ans, 0)
	}

	return ans
}
