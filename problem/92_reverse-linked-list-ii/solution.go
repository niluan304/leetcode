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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var (
		ans = &ListNode{Next: head}

		root = &ListNode{
			Val:  0,
			Next: nil,
		}
		end, start *ListNode

		count = 0
	)

	for head != nil {
		count++
		node := head     // 拷贝指针
		head = head.Next // 移动指针

		if count < left {
			start = node
			continue
		}
		if count == left {
			end = node
		}

		node.Next = root.Next
		root.Next = node

		if count == right {
			end.Next = head

			if start == nil {
				return root.Next
			}
			start.Next = root.Next
			break
		}
	}

	return ans.Next
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var (
		start = new(ListNode)
		end   *ListNode
		ans   = &ListNode{Next: head}
		count = 0
	)

	for head != nil {
		count++
		node := head     // 拷贝指针
		head = head.Next // 移动指针

		if count < left {
			start = node
			continue
		}
		if count == left {
			end = node
		}

		node.Next = start.Next
		start.Next = node

		if count == right {
			end.Next = head

			if ans.Next == end {
				return node
			}
			return ans.Next
		}
	}

	return ans.Next
}

func reverseBetween3(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var (
		start = new(ListNode)
		end   *ListNode
		ans   = head
		count = 0
	)

	for head != nil {
		count++
		node := head     // 拷贝指针
		head = head.Next // 移动指针

		if count < left {
			start = node
			continue
		}
		if count == left {
			end = node
		}

		node.Next = start.Next
		start.Next = node

		if count == right {
			end.Next = head

			if ans == end {
				return node
			}

			return ans
		}
	}

	return ans
}

func reverseBetween4(head *ListNode, left int, right int) *ListNode {
	var (
		dummy = &ListNode{Next: head} // 哨兵结点
		p0    = dummy                 // 下一个元素为反转链表的第一个元素
	)

	for i := 1; i < left; i++ {
		p0 = p0.Next
	}

	var (
		pre *ListNode // 反转后的第一个元素
		cur = p0.Next // 指针移动位置
	)

	for i := left; i < right+1; i++ {
		node := cur.Next
		cur.Next = pre

		pre = cur
		cur = node
	}

	// p0.Next： 反转前为第一个元素, 反转后为最后一个元素
	p0.Next.Next = cur

	// p0： 下一个元素为反转链表的第一个元素
	p0.Next = pre
	return dummy.Next
}
