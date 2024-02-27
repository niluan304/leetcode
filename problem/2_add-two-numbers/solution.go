package main

import (
	"math/big"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Deprecated: 链表长度最大100, 其值可以大于 math.MaxInt64
// 时间复杂度O(m+n)
// 空间复杂度O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := ToInt(l1) + ToInt(l2)
	if n == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	var (
		p0  = &ListNode{}
		cur = p0
	)

	for n > 0 {
		cur.Next = &ListNode{Val: n % 10, Next: nil}
		cur = cur.Next

		n /= 10
	}

	return p0.Next
}

func ToInt(head *ListNode) int {
	n := 0
	for head != nil {
		n = n*10 + head.Val
		head = head.Next
	}
	return n
}

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummy = &ListNode{} // 哨兵结点/虚拟头节点（构建新链表时的常用技巧）
		cur   = dummy       // 指针 cur 负责构建新链表
		carry = 0           // 记录进位
	)

	// 开始执行加法，两条链表走完且没有进位时才能结束循环
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // 记录当前节点的和
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 处理进位情况
		carry = sum / 10

		// 构建新节点
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		cur.Next = node
		cur = cur.Next
	}

	// 返回结果链表的头结点（去除虚拟头结点）
	return dummy.Next
}

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head = &ListNode{}
		cur  = head
		flag = 0
	)

	for l1 != nil && l2 != nil {
		n := l1.Val + l2.Val + flag
		flag = n / 10

		l1 = l1.Next
		l2 = l2.Next

		node := &ListNode{
			Val:  n % 10,
			Next: nil,
		}
		cur.Next = node
		cur = node
	}

	if l1 == nil {
		l1 = l2
	}

	pre := cur
	pre.Next = l1
	cur = pre.Next
	for flag > 0 {
		if cur == nil {
			pre.Next = &ListNode{}
			cur = pre.Next
		}

		n := cur.Val + flag
		cur.Val = n % 10
		flag = n / 10

		pre = cur
		cur = pre.Next
	}

	return head.Next
}

// 将链表转换为 big.Int
// - 时间复杂度：O(n + m)。
// - 空间复杂度：O(n + m)。
func addTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := ToBigInt(l1), ToBigInt(l2)

	s := num1.Add(num1, num2).String()
	dummy := &ListNode{}
	for _, b := range s {
		dummy.Val = int(b - '0')
		dummy = &ListNode{ // 头插法, dummy.Next 始终为头结点
			Val:  0,
			Next: dummy,
		}
	}
	return dummy.Next
}

func ToBigInt(root *ListNode) *big.Int {
	e := big.NewInt(1) // 表示 10^i
	num := big.NewInt(0)
	for node := root; node != nil; node = node.Next {
		x := big.NewInt(int64(node.Val))
		num = num.Add(num, x.Mul(x, e))
		e = e.Mul(e, big.NewInt(10))
	}
	return num
}
