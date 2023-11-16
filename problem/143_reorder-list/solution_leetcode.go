package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法一：线性表
// 因为链表不支持下标访问，所以我们无法随机访问链表中任意位置的元素。
//
// 因此比较容易想到的一个方法是，我们利用线性表存储该链表，然后利用线性表可以下标访问的特点，直接按顺序访问指定元素，重建该链表即可。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/reorder-list/solutions/452867/zhong-pai-lian-biao-by-leetcode-solution/
func leetcode1(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// 方法二：寻找链表中点 + 链表逆序 + 合并链表
// 注意到目标链表即为将原链表的左半端和反转后的右半端合并后的结果。
//
// 这样我们的任务即可划分为三步：
//
// 1. 找到原链表的中点（参考「876. 链表的中间结点」）。
// 我们可以使用快慢指针来 O(N) 地找到链表的中间节点。
//
// 2. 将原链表的右半端反转（参考「206. 反转链表」）。
// 我们可以使用迭代法实现链表的反转。
//
// 3. 将原链表的两端合并。
// 因为两链表长度相差不超过 1，因此直接合并即可。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/reorder-list/solutions/452867/zhong-pai-lian-biao-by-leetcode-solution/
func leetcode2(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}
