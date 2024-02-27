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
// 空间复杂度O(m+n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := ToInt(l1) + ToInt(l2)
	if n == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	var (
		list []int

		head = &ListNode{}
		node = head
	)
	for n > 0 {
		list = append(list, n%10)
		n /= 10
	}

	for i := len(list) - 1; i >= 0; i-- {
		node.Next = &ListNode{Val: list[i]}
		node = node.Next
	}
	return head.Next
}

func ToInt(head *ListNode) int {
	n := 0
	for head != nil {
		n = n*10 + head.Val
		head = head.Next
	}
	return n
}

// Deprecated: 链表长度最大100, 其值可以大于 math.MaxInt64
// 时间复杂度O(m+n)
// 空间复杂度O(1)
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	n := ToInt(l1) + ToInt(l2)
	if n == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	var cur *ListNode

	for n > 0 {
		// 反转链表, cur永远为最高位
		next := &ListNode{
			Val:  n % 10,
			Next: cur,
		}
		cur = next

		n /= 10
	}
	return cur
}

// 翻转链表
// 时间复杂度O(m+n)
// 空间复杂度O(1)
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		// 反转链表, 将问题转化为 [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/)
		arr1, arr2 = Reverse(l1), Reverse(l2)

		cur   *ListNode // 头插法的链表头
		carry = 0       // 记录进位标志
	)

	for arr1 != nil || arr2 != nil || carry > 0 {
		sum := carry // 记录 arr1.Val + arr2.Val + carry 之和
		if arr1 != nil {
			sum += arr1.Val
			arr1 = arr1.Next
		}
		if arr2 != nil {
			sum += arr2.Val
			arr2 = arr2.Next
		}

		carry = sum / 10 // 更新进位标志

		// 头插法, cur始终为头结点
		node := &ListNode{
			Val:  sum % 10,
			Next: cur,
		}
		cur = node
	}

	return cur
}

// Reverse 反转链表
func Reverse(root *ListNode) *ListNode {
	cur := root
	var pre *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 链表转数组, 从尾到头遍历数组
// 时间复杂度O(m+n)
// 空间复杂度O(n)
func addTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		arr1, arr2 = ToSlice(l1), ToSlice(l2) // 链表转数组, 向后遍历数组, 原理同栈
		n1, n2     = len(arr1), len(arr2)     // 数组当前索引位置

		cur   *ListNode // 头插法的链表头
		carry = 0       // 记录进位标志
	)

	for n1 > 0 || n2 > 0 || carry > 0 {
		sum := carry // 记录 arr1[n1] + arr2[n2] + carry 之和
		if n1 > 0 {
			n1--
			sum += arr1[n1]
		}
		if n2 > 0 {
			n2--
			sum += arr2[n2]
		}
		carry = sum / 10 // 更新进位标志

		// 头插法, cur始终为头结点
		node := &ListNode{
			Val:  sum % 10,
			Next: cur,
		}
		cur = node
	}

	return cur
}

// ToSlice 链表转数组
func ToSlice(head *ListNode) []int {
	list := make([]int, 0, 100)

	cur := head
	for cur != nil {
		list = append(list, cur.Val) // append当前节点值
		cur = cur.Next               // 下一个结点
	}
	return list
}

// 将链表转换为 big.Int
// - 时间复杂度：O(n + m)。
// - 空间复杂度：O(n + m)。
func addTwoNumbers5(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := ToBigInt(l1), ToBigInt(l2)

	s := num1.Add(num1, num2).String()

	dummy := &ListNode{}
	cur := dummy
	for _, b := range s {
		cur.Next = &ListNode{Val: int(b - '0')} // 尾插法
		cur = cur.Next
	}
	return dummy.Next
}

func ToBigInt(root *ListNode) *big.Int {
	num := big.NewInt(0)
	ten := big.NewInt(10)
	for node := root; node != nil; node = node.Next {
		x := big.NewInt(int64(node.Val))
		num = num.Add(num.Mul(num, ten), x)
	}
	return num
}
